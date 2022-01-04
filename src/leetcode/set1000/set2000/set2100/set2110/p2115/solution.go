package p2115

func maximumInvitations(favorite []int) int {
	n := len(favorite)

	g := NewGraph(n, n+1)

	deg := make([]int, n)

	for i := 0; i < n; i++ {
		g.AddEdge(favorite[i], i)
		deg[favorite[i]]++
	}

	q := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := favorite[v]

		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	var rdfs func(v int) int

	rdfs = func(v int) int {
		depth := 1
		for i := g.nodes[v]; i > 0; i = g.next[i] {
			w := g.to[i]
			if deg[w] == 0 {
				depth = max(depth, rdfs(w)+1)
			}
		}
		return depth
	}

	var maxRingSize, sumChainSize int

	for i, d := range deg {
		if d <= 0 {
			continue
		}
		deg[i] = -1
		ringSize := 1

		for v := favorite[i]; v != i; v = favorite[v] {
			deg[v] = -1
			ringSize++
		}

		if ringSize == 2 {
			sumChainSize += rdfs(i) + rdfs(favorite[i])
		} else {
			maxRingSize = max(maxRingSize, ringSize)
		}
	}

	return max(maxRingSize, sumChainSize)
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
