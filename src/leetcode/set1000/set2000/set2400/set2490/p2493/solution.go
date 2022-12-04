package p2493

func magnificentSets(n int, edges [][]int) int {
	// n <= 500
	g := NewGraph(n, len(edges)*2+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var comp []int
	color := make([]int, n)

	var dfs func(x int, c int) bool
	dfs = func(x int, c int) bool {
		if color[x] == c {
			return true
		}
		if color[x] == -c {
			return false
		}
		comp = append(comp, x)
		color[x] = c

		for i := g.nodes[x]; i > 0; i = g.next[i] {
			y := g.to[i]
			if !dfs(y, -c) {
				return false
			}
		}
		return true
	}

	var comps [][]int
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			comp = make([]int, 0, 1)
			if !dfs(i, 1) {
				return -1
			}
			comps = append(comps, comp)
		}
	}

	// then bfs
	que := make([]int, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	bfs := func(x int) int {
		var front, end int
		dist[x] = 0
		que[end] = x
		end++
		var res int
		for front < end {
			u := que[front]
			front++
			res = max(res, dist[u])
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return res + 1
	}

	maxGroup := func(comp []int) int {
		var ans int
		for _, u := range comp {
			for _, i := range comp {
				dist[i] = -1
			}
			ans = max(ans, bfs(u))
		}
		return ans
	}

	var res int
	for _, comp := range comps {
		res += maxGroup(comp)
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
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
