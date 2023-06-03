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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

const N = 200010

var F []int

func init() {
	F = append(F, 1)
	F = append(F, 1)

	for F[len(F)-1] < N {
		F = append(F, F[len(F)-2]+F[len(F)-1])
	}
}

func solve(n int, E [][]int) bool {
	it := sort.SearchInts(F, n)
	if it == len(F) || F[it] != n {
		return false
	}
	g := NewGraph(n, 2*n)

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}
	sz := make([]int, n)
	var dfs1 func(p int, u int)
	fa := make([]int, n)
	dfs1 = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				fa[v] = u
				dfs1(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs1(0, 0)

	marked := make([]bool, n)

	var mark func(p int, u int, it int, res *int)

	mark = func(p int, u int, it int, res *int) {
		if *res == -1 {
			if sz[u] == F[it-1] || sz[u] == F[it-2] {
				*res = u
			}
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if v != p && !marked[v] {
					mark(u, v, it, res)
				}
			}
		}

	}

	var dfs2 func(p int, u int, it int) bool

	dfs2 = func(p int, u int, it int) bool {
		if it < 2 {
			return true
		}

		v := -1

		mark(p, u, it, &v)

		if v < 0 {
			return false
		}

		x := fa[v]
		for x != u {
			sz[x] -= sz[v]
			x = fa[x]
		}
		sz[u] -= sz[v]

		marked[v] = true

		if sz[v] == F[it-1] {
			return dfs2(fa[v], v, it-1) && dfs2(p, u, it-2)
		}
		return dfs2(fa[v], v, it-2) && dfs2(p, u, it-1)
	}

	return dfs2(0, 0, it)
}
