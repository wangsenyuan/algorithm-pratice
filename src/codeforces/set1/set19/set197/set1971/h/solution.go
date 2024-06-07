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
		n := readNum(reader)
		grid := make([][]int, 3)
		for i := 0; i < 3; i++ {
			grid[i] = readNNums(reader, n)
		}
		res := solve(grid)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(grid [][]int) bool {
	n := len(grid[0])
	N := n + 1

	g := NewGraph(N*2, 10*n)

	connect := func(u, v int) {
		g.AddEdge(-u+N, v+N)
		g.AddEdge(-v+N, u+N)
	}

	for i := 0; i < n; i++ {
		a, b, c := grid[0][i], grid[1][i], grid[2][i]
		connect(a, b)
		connect(b, c)
		connect(c, a)
	}
	dfn := make([]int, 2*N)

	low := make([]int, 2*N)

	var cnt int
	var id int

	var tarjan func(u int)

	stack := make([]int, 2*N+1)
	var top int

	color := make([]int, 2*N)

	vis := make([]bool, 2*N)

	tarjan = func(u int) {
		cnt++
		low[u] = cnt
		dfn[u] = cnt
		vis[u] = true

		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dfn[v] == 0 {
				tarjan(v)
				low[u] = min(low[u], low[v])
			} else if vis[v] {
				low[u] = min(low[u], dfn[v])
			}
		}

		if dfn[u] == low[u] {
			id++
			for top > 0 {
				v := stack[top-1]
				color[v] = id
				vis[v] = false
				top--
				if v == u {
					break
				}
			}
		}
	}

	for i := 1; i <= 2*n; i++ {
		if dfn[i] == 0 {
			tarjan(i)
		}
	}

	for i := 1; i <= n; i++ {
		if color[N+i] == color[N-i] {
			return false
		}
	}
	return true
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
