package p1857

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	g := NewGraph(n, len(edges))

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}
	vis := make([]int, n)
	var dfs func(u int) bool

	stack := make([]int, 0, n)

	dfs = func(u int) bool {
		if vis[u] == 1 {
			// a loop
			return false
		}
		vis[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] < 2 {
				if !dfs(v) {
					return false
				}
			}
		}
		stack = append(stack, u)
		vis[u]++
		return true
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if !dfs(i) {
				return -1
			}
		}
	}
	var best int
	// stack is the topo order
	count := make([][]int, n)

	for i := 0; i < n; i++ {
		count[i] = make([]int, 26)
	}

	for i := n - 1; i >= 0; i-- {
		u := stack[i]
		count[u][int(colors[u]-'a')]++

		for j := g.nodes[u]; j > 0; j = g.next[j] {
			v := g.to[j]

			for k := 0; k < 26; k++ {
				count[v][k] = max(count[u][k], count[v][k])
			}
		}
		for k := 0; k < 26; k++ {
			best = max(best, count[u][k])
		}
	}

	return best
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

func NewGraph(n, e int) Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	return Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
