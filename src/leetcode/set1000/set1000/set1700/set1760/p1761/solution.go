package p1761

func minTrioDegree(n int, edges [][]int) int {
	g := NewGraph(n, len(edges)*2+10)
	degree := make([]int, n)
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		u--
		v--
		grid[u][v] = true
		grid[v][u] = true
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		degree[u]++
		degree[v]++
	}
	res := -1
	for u := 0; u < n; u++ {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			for j := g.next[i]; j > 0; j = g.next[j] {
				w := g.to[j]
				if grid[v][w] {
					tmp := degree[w] + degree[u] + degree[v]
					tmp -= 6
					if res < 0 || res > tmp {
						res = tmp
					}
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
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
