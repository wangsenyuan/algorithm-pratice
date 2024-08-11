package p3249

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	var dfs func(p int, u int)
	var res int
	dfs = func(p int, u int) {
		sz[u] = 1
		cz := -1
		ok := true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				if cz < 0 {
					cz = sz[v]
				} else if ok {
					ok = cz == sz[v]
				}
			}
		}

		if ok {
			res++
		}
	}
	dfs(-1, 0)

	return res
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
