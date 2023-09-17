package p2859

func minEdgeReversals(n int, edges [][]int) []int {
	// len(edges) == n - 1
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1], 1)
		g.AddEdge(e[1], e[0], -1)
	}

	sz := make([]int, n)
	// 在子树u中需要反转的次数
	dp := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				// 如果是从u到v的，不需要反转，否则需要反转
				dp[u] += dp[v]
				if g.val[i] < 0 {
					dp[u]++
				}
			}
		}
	}
	dfs(0, 0)
	ans := make([]int, n)
	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		ans[u] = dp[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dp[u]
				// cancel out v contribution
				dp[u] -= dp[v]
				if g.val[i] < 0 {
					dp[u]--
				}
				// contribute u to v
				dp[v] += dp[u]
				if g.val[i] > 0 {
					dp[v]++
				}
				dfs2(u, v)
				// restore u, no need for v
				dp[u] = tmp
			}
		}
	}

	dfs2(0, 0)
	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
