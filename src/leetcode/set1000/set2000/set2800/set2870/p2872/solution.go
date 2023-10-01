package p2872

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int) int
	var res int

	dfs = func(p int, u int) int {
		cur := values[u] % k
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cur += dfs(u, v)
				cur %= k
			}
		}
		if cur == 0 {
			res++
		}
		return cur
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
