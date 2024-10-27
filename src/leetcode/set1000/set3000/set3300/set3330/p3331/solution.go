package p3331

func findSubtreeSizes(parent []int, s string) []int {

	n := len(parent)
	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		g.AddEdge(parent[i], i)
	}

	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		last[i] = -1
	}

	g2 := NewGraph(n, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		x := int(s[u] - 'a')
		if last[x] != -1 {
			g2.AddEdge(last[x], u)
		} else if p >= 0 {
			g2.AddEdge(p, u)
		}
		ov := last[x]
		last[x] = u

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(u, v)
		}

		last[x] = ov
	}

	dfs(-1, 0)

	sz := make([]int, n)

	var dfs2 func(u int)
	dfs2 = func(u int) {
		sz[u]++
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			dfs2(v)
			sz[u] += sz[v]
		}
	}

	dfs2(0)

	return sz
}

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
