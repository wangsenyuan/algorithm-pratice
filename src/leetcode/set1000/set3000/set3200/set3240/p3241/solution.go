package p3241

func timeTaken(edges [][]int) []int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	type pair struct {
		first  int
		second int
	}

	res := make([][]pair, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		res[u] = []pair{{0, u}, {0, u}}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				tmp := res[v][0].first + 2 - v%2
				if tmp >= res[u][0].first {
					res[u][1] = res[u][0]
					res[u][0] = pair{tmp, v}
				} else if tmp >= res[u][1].first {
					res[u][1] = pair{tmp, v}
				}
			}
		}
	}

	dfs(-1, 0)

	ans := make([]int, n)

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		// 如果以u为root
		if p >= 0 {
			x := res[p][0].first
			if res[p][0].second == u {
				x = res[p][1].first
			}
			x += 2 - p%2
			if x >= res[u][0].first {
				res[u][1] = res[u][0]
				res[u][0] = pair{x, p}
			} else if x >= res[u][1].first {
				res[u][1] = pair{x, p}
			}
		}
		ans[u] = res[u][0].first
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)

	return ans
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
