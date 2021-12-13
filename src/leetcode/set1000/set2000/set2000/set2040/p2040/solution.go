package p2040

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)

	g := NewGraph(n, len(edges)*2+2)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	que := make([]int, n)
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = -1
	}

	var front, end int
	que[end] = 0
	end++
	for front < end {
		u := que[front]
		front++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				que[end] = v
				end++
				dist[v] = dist[u] + 1
			}
		}
	}

	var res int

	for i := 1; i < n; i++ {
		x := 2 * dist[i]
		if x <= patience[i] {
			res = max(res, x+1)
			continue
		}
		// i will reset messages to master until x
		d := (x + patience[i] - 1) / patience[i]
		y := (d-1)*patience[i] + x + 1
		res = max(res, y)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
