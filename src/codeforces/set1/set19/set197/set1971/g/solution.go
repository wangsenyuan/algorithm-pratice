package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

const H = 30

func solve(a []int) []int {

	var groups [][]int

	n := len(a)

	belong := make([]int, n)

	tr := NewTrie()

	for i, num := range a {
		j := GetVal(tr, num)
		if j >= 0 {
			g := belong[j]
			groups[g] = append(groups[g], i)
		} else {
			belong[i] = len(groups)
			groups = append(groups, []int{i})
			AddNum(tr, num, i)
		}
	}

	res := make([]int, n)

	fill := func(arr []int) {
		val := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			val[i] = a[arr[i]]
		}
		sort.Ints(val)

		for i := 0; i < len(arr); i++ {
			res[arr[i]] = val[i]
		}
	}

	for _, g := range groups {
		fill(g)
	}

	return res
}

type Trie struct {
	children []*Trie
	val      int
}

func NewTrie() *Trie {
	tr := new(Trie)
	tr.children = make([]*Trie, 2)
	return tr
}

func AddNum(root *Trie, num int, id int) {
	cur := root
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if cur.children[x] == nil {
			cur.children[x] = NewTrie()

		}
		cur = cur.children[x]
	}

	cur.val = id
}

func GetVal(root *Trie, num int) int {
	cur := root
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if cur.children[x] == nil {
			if i >= 2 {
				return -1
			}
			cur = cur.children[x^1]
		} else {
			cur = cur.children[x]
		}
	}
	return cur.val
}
