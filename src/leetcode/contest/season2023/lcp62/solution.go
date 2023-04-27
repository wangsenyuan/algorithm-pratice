package lcp62

func transportationHub(path [][]int) int {
	var n int
	for _, p := range path {
		n = max(n, p[0])
		n = max(n, p[1])
	}
	n++
	g := NewGraph(n, len(path))
	deg := make([]int, n)
	for _, p := range path {
		u, v := p[0], p[1]
		g.AddEdge(v, u)
		deg[u]++
	}

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			out := make(map[int]int)
			for j := g.nodes[i]; j > 0; j = g.next[j] {
				out[g.to[j]]++
			}
			if len(out) == n-1 {
				return i
			}
		}
	}

	return -1
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
