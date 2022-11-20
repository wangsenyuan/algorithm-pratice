package p2477

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1

	g := NewGraph(n, 2*n-1)

	for _, road := range roads {
		u, v := road[0], road[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	//假设考虑节点u，它的子节点的数目是sz(u),
	// 如果 sz[u] <= seats， 那么它们应该乘坐节点u的车

	var ans int64

	var dfs func(p int, u int)

	sz := make([]int, n)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
		sz[u]++
		if p >= 0 {
			ans += int64((sz[u] + seats - 1) / seats)
		}
	}

	dfs(-1, 0)

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
