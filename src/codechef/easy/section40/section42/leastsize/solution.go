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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		root, P := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", root))
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", P[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, E [][]int) (int, []int) {
	// lca(P[i], P[i+1]) 组成的set的size最小
	// 如果先找到重心点，这个就是root

	g := createGraphFromEdges(n, E)
	sz := make([]int, n)
	big := make([]int, n)

	fa := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = p
		big[u] = -1
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[v] > sz[big[u]] {
					big[u] = v
				}
			}
		}
	}

	dfs(-1, 0)

	var root int

	for big[root] > 0 && sz[big[root]]*2 >= n {
		root = big[root]
	}

	// fa[root] >= 0
	if big[root] < 0 || n-sz[root] > sz[big[root]] {
		// big of root change to fa
		big[root] = fa[root]
	}

	que := make([]int, n)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	dist[root] = 0

	bfs := func(x int) []int {
		dist[x] = 1
		var front, end int
		que[end] = x
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					que[end] = v
					end++
					dist[v] = dist[u] + 1
				}
			}
		}

		return que[:end]
	}
	var it int
	P := make([]int, n)
	place := func(u int) {
		nodes := bfs(u)
		for _, v := range nodes {
			P[it] = v + 1
			it += 2
			if it >= n {
				it = 1
			}
		}
	}

	if big[root] >= 0 {
		place(big[root])
	}

	for i := g.nodes[root]; i > 0; i = g.next[i] {
		v := g.to[i]
		if v != big[root] {
			place(v)
		}
	}

	P[it] = root + 1

	return root + 1, P
}

func createGraphFromEdges(n int, E [][]int) *Graph {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	return g
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
