package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		m := readNum(reader)
		maps := make([][]string, m)
		for i := 0; i < m; i++ {
			n := readNum(reader)
			edges := make([]string, n-1)
			for j := 0; j < n-1; j++ {
				edges[j] = readString(reader)
			}
			maps[i] = edges
		}
		res := solve(m, maps)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteString(fmt.Sprintln())
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(m int, maps [][]string) []int {

	var root *Node

	for i := 0; i < m; i++ {
		root = merge(root, NewTree(maps[i]))
	}

	cnt := make([]int, m+1)

	var dfs func(node *Node, d int)

	dfs = func(node *Node, d int) {
		if node == nil {
			return
		}
		cnt[node.cnt] = max(cnt[node.cnt], d)
		dfs(node.left, d+1)
		dfs(node.right, d+1)
	}

	dfs(root, 0)

	for i := 2; i < m; i++ {
		cnt[i] = min(cnt[i], cnt[i-1])
	}
	return cnt[1:]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func merge(a *Node, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	// b.cnt = 1
	a.cnt += b.cnt
	a.left = merge(a.left, b.left)
	a.right = merge(a.right, b.right)
	return a
}

type Node struct {
	left  *Node
	right *Node
	cnt   int
}

func NewTree(spec []string) *Node {
	n := len(spec) + 1
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = new(Node)
		nodes[i].cnt++
	}

	for _, cur := range spec {
		bs := []byte(cur)
		var u, v int
		pos := readInt(bs, 0, &u) + 1
		c := bs[pos]
		readInt(bs, pos+2, &v)
		u--
		v--
		if c == 'L' {
			nodes[u].left = nodes[v]
		} else {
			nodes[u].right = nodes[v]
		}
	}
	return nodes[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
