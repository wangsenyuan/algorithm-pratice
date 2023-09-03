package p2846

const H = 15

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	P := make([][]int, n)
	depth := make([]int, n)
	val := make([][]int, n)
	for i := 0; i < n; i++ {
		val[i] = make([]int, 26)
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				copy(val[v], val[u])
				depth[v] = depth[u] + 1
				val[v][g.val[i]-1]++
				dfs(u, v)
			}
		}
	}
	dfs(0, 0)

	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if depth[u]-(1<<i) >= depth[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	check := func(u, v int) int {
		p := lca(u, v)
		var ans int
		for d := 0; d < 26; d++ {
			ans = max(ans, val[u][d]+val[v][d]-2*val[p][d])
		}
		cnt := depth[u] + depth[v] - 2*depth[p]
		return cnt - ans
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur[0], cur[1])
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
