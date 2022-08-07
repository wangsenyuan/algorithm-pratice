package p2368

func reachableNodes(n int, edges [][]int, restricted []int) int {
	g := NewGraph(n, len(edges)*2)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	block := make([]bool, n)

	for _, x := range restricted {
		block[x] = true
	}

	var dfs func(p, u int) int

	dfs = func(p, u int) int {
		var res int
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !block[v] {
				res += dfs(u, v)
			}
		}
		res++
		return res
	}

	return dfs(0, 0)
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
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
