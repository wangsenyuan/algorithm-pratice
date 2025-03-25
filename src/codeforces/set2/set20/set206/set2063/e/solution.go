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
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	de := make([]int, n)
	dc := make([]int, n)
	sz := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dc[de[u]]++
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p {
				continue
			}
			de[v] = de[u] + 1
			dfs(u, v)
			sz[u] += sz[v]
		}
	}
	dfs(-1, 0)

	sc := make([]int, n+1)
	copy(sc, dc)
	for i := n - 1; i >= 0; i-- {
		sc[i] += sc[i+1]
	}

	var ans, ans2 int

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		// u is min
		// sc[d[u]]表示至少和u一样深的节点的总数量，其中包括了它自己（以及子节点）
		ans += 2 * de[u] * (sc[de[u]] - sz[u])
		// u is lca
		sub_cnt := sz[u] - 1
		l_cnt := 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p {
				continue
			}
			l_cnt += (sub_cnt - sz[v]) * sz[v]
			dfs2(u, v)
		}
		ans2 += (2*de[u] + 1) * (l_cnt / 2)
	}

	dfs2(-1, 0)

	for i := range n {
		ans2 += i * dc[i] * (dc[i] - 1)
	}

	return ans - ans2
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
