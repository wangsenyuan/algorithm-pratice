package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k) + 1
		queries[i] = make([]int, k)
		for j := 0; j < k; j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	res := solve(n, edges, queries)

	var buf bytes.Buffer

	for _, x := range res {
		if x {
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

func solve(n int, edges [][]int, queries [][]int) []bool {
	g := NewGraph(n, 2*n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	dep := make([]int, n)
	fa := make([]int, n)
	var dfs func(p int, u int)

	in := make([]int, n)
	out := make([]int, n)

	var timer int

	dfs = func(p int, u int) {
		in[u] = timer
		timer++
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
		out[u] = timer
	}

	dfs(-1, 0)

	isAncestor := func(u int, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	check := func(vs []int) bool {
		var f int
		for _, v := range vs {
			v--
			if dep[v] > dep[f] {
				f = v
			}
		}

		for _, v := range vs {
			v--
			if v == 0 || fa[v] == 0 {
				continue
			}
			if isAncestor(v, f) || isAncestor(fa[v], f) {
				continue
			}
			return false
		}

		return true
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur)
	}

	return ans
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
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
