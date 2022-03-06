package p2192

import "sort"

func getAncestors(n int, edges [][]int) [][]int {

	g := NewGraph(n, len(edges)+2)

	for _, ed := range edges {
		f, t := ed[0], ed[1]
		g.AddEdge(t, f)
	}
	vis := make([]int, n)

	var dfs func(u int, x int) int
	dfs = func(u int, x int) int {
		vis[u] = x
		cnt := 1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] != x {
				cnt += dfs(v, x)
			}
		}
		return cnt
	}
	res := make([][]int, n)
	for u := 0; u < n; u++ {
		cnt := dfs(u, u+1)
		// without
		res[u] = make([]int, 0, cnt-1)
		for v := 0; v < n; v++ {
			if u == v || vis[v] != u+1 {
				continue
			}
			res[u] = append(res[u], v)
		}
		sort.Ints(res[u])
	}

	return res
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e)
	g.to = make([]int, e)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.to[g.cur] = v
	g.node[u] = g.cur
}
