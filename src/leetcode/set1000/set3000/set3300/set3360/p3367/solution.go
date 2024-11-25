package p3367

import "sort"

func maximizeSumOfWeights(edges [][]int, k int) int64 {
	n := len(edges) + 1
	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	// dp[u][0] 表示不保留从父节点进入的边时, 保证u子树满足k约束时的最优解
	// dp[u][1] 表示必须保留进入边时的最优解
	dp := make([][2]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		var sum int
		var arr []int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if p != v {
				dfs(u, v)
				sum += dp[v][1] + w
				arr = append(arr, dp[v][0]-(dp[v][1]+w))
			}
		}
		sort.Ints(arr)
		dp[u][0] = sum
		dp[u][1] = sum
		if len(arr) > k {
			// 删除多余的边
			for i := len(arr); i > k; i-- {
				dp[u][0] += arr[i-1]
			}
		}
		tmp := dp[u][0]
		// dp[u][0]至少这么大, 但可以继续增加
		for i := min(len(arr), k); i > 0; i-- {
			tmp += arr[i-1]
			dp[u][0] = max(dp[u][0], tmp)
		}
		if len(arr) > k-1 {
			for i := len(arr); i > k-1; i-- {
				dp[u][1] += arr[i-1]
			}
		
		tmp = dp[u][1]
		for i := min(len(arr), k-1); i > 0; i-- {
			tmp += arr[i-1]
			dp[u][1] = max(dp[u][1], tmp)
		}
	}

	dfs(-1, 0)

	return int64(dp[0][0])
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
