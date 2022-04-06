package p310

func findMinHeightTrees(n int, edges [][]int) []int {
	g := NewGraph(n, 2*(n-1))

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	H := make([]int, n)

	var dfs func(p, u int)
	dfs = func(p, u int) {
		H[u] = D[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
				H[u] = max(H[u], H[v])
			}
		}
	}

	dfs(0, 0)

	var a int

	for i := 0; i < n; i++ {
		if D[a] < D[i] {
			a = i
		}
	}

	// a is one endpoint
	D[a] = 0
	dfs(a, a)

	diameter := H[a] + 1

	half := diameter / 2

	var res []int

	for i := 0; i < n; i++ {
		if H[i] == diameter-1 {
			if D[i] == half {
				res = append(res, i)
			} else if diameter%2 == 0 && D[i] == half-1 {
				res = append(res, i)
			}
		}
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
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
