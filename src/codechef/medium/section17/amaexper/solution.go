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
			E[i] = readNNums(reader, 3)
		}
		res := solve(n, E)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, E [][]int) []int {
	// strategic cost of a node u is an sub-tree rooted at x, is the farthest distance in that sub-tree from node u
	// strategic cost of a sub-tree = min stretegic cost of nodes in the sub-tree
	// 假设对于x，的所以的child，已经知道它们的cost （子树的cost）
	// 如何计算x的cost呢？
	// 对于child u来说， u的新的cost = 到其他子树的最远距离 （这个很好算）
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dmax := make([][]int, n)
	mnode := make([][]int, n)

	for i := 0; i < n; i++ {
		dmax[i] = make([]int, 2)
		mnode[i] = make([]int, 2)
	}
	radius := make([]int, n)
	droot := make([]int, n)
	diam := make([]int, n)
	parent := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		mnode[u][0] = u
		mnode[u][1] = u

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			parent[v] = u
			droot[v] = droot[u] + g.val[i]

			dfs(u, v)

			if diam[v] > diam[u] {
				diam[u] = diam[v]
				radius[u] = radius[v]
			}
			if dmax[v][0]+g.val[i] > dmax[u][0] {
				dmax[u][1] = dmax[u][0]
				mnode[u][1] = mnode[u][0]
				dmax[u][0] = dmax[v][0] + g.val[i]
				mnode[u][0] = mnode[v][0]
			} else if dmax[v][0]+g.val[i] > dmax[u][1] {
				dmax[u][1] = dmax[v][0] + g.val[i]
				mnode[u][1] = mnode[v][0]
			}
		}
		if dmax[u][0]+dmax[u][1] > diam[u] {
			diam[u] = dmax[u][0] + dmax[u][1]
			for mnode[u][0] != u && (dmax[u][1]+droot[parent[mnode[u][0]]]-droot[u] >= ((diam[u] + 1) >> 1)) {
				mnode[u][0] = parent[mnode[u][0]]
			}
			radius[u] = dmax[u][1] + droot[mnode[u][0]] - droot[u]
			if mnode[u][0] != u && (diam[u]-dmax[u][1]-(droot[parent[mnode[u][0]]]-droot[u])) < radius[u] {
				radius[u] = diam[u] - dmax[u][1] - (droot[parent[mnode[u][0]]] - droot[u])
			}
		}
	}

	dfs(0, 0)

	return radius
}

func min(a, b int) int {
	return a + b - max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
