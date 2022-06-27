package p2316

func countPairs(n int, edges [][]int) int64 {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	var dfs func(u int) int

	dfs = func(u int) int {
		var res int
		vis[u] = true
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				res += dfs(v)
			}
		}
		return res + 1
	}

	var res int64

	for i := 0; i < n; i++ {
		if !vis[i] {
			m := dfs(i)
			res += int64(m) * int64(n-m)
		}
	}

	return res / 2
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
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
