package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Trie struct {
	children [2]*Trie
	val      int
}

const D = 30

func (root *Trie) Add(num int, v int) {
	node := root
	for i := D - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if node.children[x] == nil {
			node.children[x] = new(Trie)
		}
		node = node.children[x]
	}
	node.val = v
}

func (root *Trie) Find(num int) int {
	node := root
	for i := D - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if node.children[x] == nil {
			return -1
		}
		node = node.children[x]
	}
	return node.val
}

func solve(nums []int, queries [][]int) []bool {
	// 如果 xor[l...r] = 0 => yes
	// else x = xor[l...r] 那么，如果存在某个 xor[i..r] = x => yes
	// xor[r] ^ xor[i-1] = x
	// 找最近的那个位置就可以了
	n := len(nums)
	ans := make([]bool, len(queries))

	xor := make([]int, n+1)
	for r, num := range nums {
		xor[r+1] = xor[r] ^ num
	}

	begin := make([][]int, n)
	end := make([][]int, n)

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		val := xor[r] ^ xor[l-1]
		if val == 0 {
			ans[i] = l < r
			continue
		}
		begin[l-1] = append(begin[l-1], i)
		end[r-1] = append(end[r-1], i)
	}

	mid := make([]int, len(queries))

	tr := new(Trie)
	tr.Add(0, 0)

	for r := 0; r < n; r++ {
		for _, i := range end[r] {
			l := queries[i][0] - 1
			j := tr.Find(xor[l])
			if l < j {
				mid[i] = j
			}
		}
		tr.Add(xor[r+1], r+1)
	}

	xor[n] = 0

	tr = new(Trie)
	tr.Add(0, n)

	for l := n - 1; l >= 0; l-- {
		xor[l] = xor[l+1] ^ nums[l]
		for _, i := range begin[l] {
			r := queries[i][1] - 1
			j := tr.Find(xor[r+1])
			if j < r && mid[i] > j {
				ans[i] = true
			}
		}
		tr.Add(xor[l], l)
	}

	return ans
}
