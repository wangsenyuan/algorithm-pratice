package p2924

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	// dp[u] = 在保证子树u健康的情况下的，最大的value的和
	// sum[u] = 子树的和
	// dp[v] = sum of (dp[u]) max sum[v] - value[v]
	n := len(values)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	dp := make([]int64, n)
	sum := make([]int64, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		var tmp int64
		leaf := true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sum[u] += sum[v]
				tmp += dp[v]
				leaf = false
			}
		}
		if !leaf {
			dp[u] = max(tmp+int64(values[u]), sum[u])
		}
		// else dp[u] = 0
		sum[u] += int64(values[u])
	}
	dfs(0, 0)

	return dp[0]
}

func max(a, b int64) int64 {
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

func NewGraph(n, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+10)
	g.to = make([]int, e+10)
	// start from 2, reverse of edge i is i ^ 1
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
