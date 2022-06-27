package p2322

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)

	X := make([]int, n)

	var dfs func(p, u int, t *int)

	dfs = func(p, u int, t *int) {
		X[u] = nums[u]
		in[u] = *t
		*t++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, t)
				X[u] ^= X[v]
			}
		}
		out[u] = *t
	}

	dfs(0, 0, new(int))

	isAncester := func(u, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	best := 1 << 30

	for u := 1; u < n; u++ {
		for v := u + 1; v < n; v++ {
			// remove edge u, v
			a := X[0]
			var b, c int
			if isAncester(u, v) {
				a ^= X[u]
				b = X[u] ^ X[v]
				c = X[v]
			} else if isAncester(v, u) {
				a ^= X[v]
				b = X[v] ^ X[u]
				c = X[u]
			} else {
				a ^= X[u] ^ X[v]
				b = X[u]
				c = X[v]
			}

			best = min(best, max3(a, b, c)-min3(a, b, c))
		}
	}

	return best
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
