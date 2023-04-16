package p2646

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	// 先计算出各个trip的lca
	// 把trip都拆分成从u到lca的部分 （或者lca的某个子节点）
	// 然后进行dp
	// dp[u][0] 表示u（u节点不半价)的子树中，trip的price之和
	// dp[u][1] 表示半价时的trip之和
	// 现在考虑u的父节点p
	// 如果 p进行了半价, 那么 dp[p][1] = sum(dp[u][0]) for children of p
	//  + 以p为lca的trip fp[u][0]
	// fp[u][0]表示某个trip的端点在u的子树中，且u不进行半价时的最小sum
	// fp[u][0] = price[u] + sum min(fp[v][0], fp[v][1])
	// fp[u][1] = price[u] / 2 + sum fp[v][0]

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	D := make([]int, n)
	P := make([]int, n)
	big := make([]int, n)

	be := make([]int, n)

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		P[u] = p
		sz[u]++
		big[u] = -1

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			D[v] = 1 + D[u]
			dfs1(u, v)
			sz[u] += sz[v]
			if big[u] < 0 || sz[big[u]] < sz[v] {
				big[u] = v
			}
		}
	}

	dfs1(-1, 0)

	var dfs2 func(p int, u int, x int)

	dfs2 = func(p int, u int, x int) {
		be[u] = x
		if big[u] >= 0 {
			dfs2(u, big[u], x)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && v != big[u] {
				dfs2(u, v, v)
			}
		}
	}

	dfs2(-1, 0, 0)

	lca := func(u, v int) int {
		for be[u] != be[v] {
			if D[be[u]] < D[be[v]] {
				u, v = v, u
			}
			u = P[be[u]]
		}
		if D[u] > D[v] {
			u, v = v, u
		}
		return u
	}

	cnt := make([]int, n)

	for _, trip := range trips {
		u, v := trip[0], trip[1]
		if D[u] > D[v] {
			u, v = v, u
		}
		p := lca(u, v)
		if p == v {
			cnt[u]++
			if p != 0 {
				cnt[P[p]]--
			}
		} else {
			// u 到 p
			// v 到 pv
			cnt[u]++
			if P[p] >= 0 {
				cnt[P[p]]--
			}

			cnt[v]++
			cnt[p]--
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = 1 << 30
		}
	}
	var dfs3 func(p int, u int)

	dfs3 = func(p int, u int) {

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs3(u, v)
				cnt[u] += cnt[v]
			}
		}
		// dp[u][0] = 不半件
		// dp[u][1] = 半价
		dp[u][0] = cnt[u] * price[u]
		dp[u][1] = cnt[u] * price[u] / 2
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dp[u][0] += min(dp[v][0], dp[v][1])
				dp[u][1] += dp[v][0]
			}
		}
	}

	dfs3(-1, 0)

	return min(dp[0][0], dp[0][1])
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
