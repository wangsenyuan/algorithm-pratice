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
		routes := make([][]int, m)
		for i := 0; i < m; i++ {
			routes[i] = readNNums(reader, 2)
		}
		k, res := solve(n, routes)
		buf.WriteString(fmt.Sprintf("%d\n", k))
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, routes [][]int) (int, []int) {

	deg := make([]int, n)

	for _, r := range routes {
		u, v := r[0]-1, r[1]-1
		deg[u]++
		deg[v]++
	}

	for i := 0; i < n; i++ {
		if deg[i] < n-1 {
			return solve1(routes, i, n)
		}
	}
	m := len(routes)
	ans := make([]int, m)
	color := 2
	for i, r := range routes {
		u, v := r[0]-1, r[1]-1
		ans[i] = 1
		if u == 0 || v == 0 {
			ans[i] = color
			color = 5 - color
		}
	}
	return 3, ans
}

func solve1(routes [][]int, x int, n int) (int, []int) {
	// u的所有的边都是颜色1，其他的边都是颜色2
	m := len(routes)
	ans := make([]int, m)
	for i, r := range routes {
		u, v := r[0]-1, r[1]-1
		ans[i] = 1
		if x == u || x == v {
			ans[i] = 2
		}
	}

	return 2, ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
