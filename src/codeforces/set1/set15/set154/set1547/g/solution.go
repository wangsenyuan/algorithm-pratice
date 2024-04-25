package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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

const oo = 1 << 60

func solve(n int, routes [][]int) []int {

	g := NewGraph(n, len(routes))

	for _, route := range routes {
		u, v := route[0]-1, route[1]-1
		g.AddEdge(u, v)
	}

	vis := make([]bool, n)
	var dfs func(u int, cycle bool)

	res := make([]int, n)

	dfs = func(u int, cycle bool) {
		vis[u] = true

		if cycle {
			res[u] = -1
		} else {
			res[u]++
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if res[v] < 0 {
				continue
			}
			if cycle || vis[v] {
				// w 在环上，再访问一次（w 至多访问三次）
				// 从 w 出发能到达的点都在环上
				dfs(v, true)
			} else if res[v] < 2 {
				// 除非后面发现 w 在环上，否则至多访问 w 两次
				dfs(v, false)
			}
		}
		vis[u] = false
	}

	dfs(0, false)

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)

		routes := make([][]int, m)
		for i := 0; i < m; i++ {
			routes[i] = readNNums(reader, 2)
		}

		res := solve(n, routes)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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
