package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		ss := make([]string, n)
		for i := 0; i < n; i++ {
			bs, _ := scanner.ReadBytes('\n')
			ss[i] = string(bs[len(bs)-1])
		}

		fmt.Println(solve(n, ss))
	}
}

var P []bool

const MAX_N = 1000007

const MOD = 1000000000007
const D = 26

var PW []int64

var PREF, SUF []int64

func init() {
	P = make([]bool, MAX_N)
	PREF = make([]int64, MAX_N)
	SUF = make([]int64, MAX_N)
	PW = make([]int64, MAX_N)

	PW[0] = 1
	for i := 1; i < MAX_N; i++ {
		PW[i] = (PW[i-1] * D) % MOD
	}
}

func reset() {
	for i := 0; i < MAX_N; i++ {
		P[i] = false
	}
}

func solve(n int, ss []string) int64 {
	reset()

	var res int64

	get := func(s string, trie *Node) int64 {
		n := len(s)
		for i := 0; i <= n; i++ {
			P[i] = false
		}
		checkPalindrome(s, P)
		r := reverse(s)
		i := -1

		var res int64
		node := trie

		for {
			if node.children[int(r[i+1]-'a')] == nil {
				break
			}
			i++
			node = node.children[int(r[i]-'a')]
			if P[i+1] {
				res += int64(node.up)
			}

			if i == n-1 {
				res += int64(node.down)
				break
			}
		}
		return res
	}

	insert := func(s string, trie *Node) {
		n := len(s)
		for i := 0; i <= n; i++ {
			P[i] = false
		}

		checkPalindrome(reverse(s), P)

		node := trie
		i := -1

		for {
			x := int(s[i+1] - 'a')
			if node.children[x] == nil {
				node.children[x] = NewNode()
			}
			i++
			node = node.children[x]
			if i == n-1 {
				node.up++
				break
			}
			if P[i+1] {
				node.down++
			}
		}
	}

	a := NewNode()

	for i := 0; i < n; i++ {
		res += get(ss[i], a)
		insert(ss[i], a)
	}

	b := NewNode()

	for i := n - 1; i >= 0; i-- {
		res += get(ss[i], b)
		insert(ss[i], b)
	}

	return res
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

type Node struct {
	children []*Node
	up       int
	down     int
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, D)
	node.up = 0
	node.down = 0
	return node
}

func getHash(s []int64, l, r int) int64 {
	return (s[r+1] - (s[l]*PW[r-l+1])%MOD + MOD) % MOD
}

func checkPalindrome(s string, res []bool) {
	n := len(s)

	PREF[0] = 0
	SUF[0] = 0

	for j := 0; j < n; j++ {
		PREF[j+1] = (PREF[j]*D + int64(s[j]-'a')) % MOD
		SUF[j+1] = (SUF[j]*D + int64(s[n-j-1]-'a')) % MOD
	}

	check := func(k int) bool {
		for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for j := n - 1; j > 0; j-- {
		a := getHash(PREF, 0, j-1)
		b := getHash(SUF, n-j, n-1)
		if a == b && check(j) {
			res[n-j] = true
		}
	}

	res[n] = true
}
