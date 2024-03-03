package p3070

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)

	f0, f1 := 0, -(1 << 60)

	for i := 0; i < n; i++ {
		f0, f1 = max(f0+nums[i], f1+(nums[i]^k)), max(f0+(nums[i]^k), f1+nums[i])
	}

	return int64(f0)
}

func maximumValueSum1(nums []int, k int, edges [][]int) int64 {
	n := len(nums)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	// dp[u][0] 表示u的入边被翻转/未翻转时的最大值
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		dp[u][1] = -(1 << 60)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				r0 := max(dp[v][0]+nums[v], dp[v][1]+(nums[v]^k))
				r1 := max(dp[v][1]+nums[v], dp[v][0]+(nums[v]^k))
				a, b := dp[u][0], dp[u][1]
				dp[u][0] = max(a+r0, b+r1)
				dp[u][1] = max(a+r1, b+r0)
			}
		}

	}
	dfs(-1, 0)

	return int64(max(dp[0][0]+nums[0], dp[0][1]+(nums[0]^k)))
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
