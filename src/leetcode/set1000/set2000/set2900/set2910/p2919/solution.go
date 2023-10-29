package p2919

const H = 20

const oo = 1 << 30

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}
	// dp[u][i] = 在处理u节点，且u的父节点使用了操作2时的最大值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, H)
		for j := 0; j < H; j++ {
			dp[i][j] = -oo
		}
	}

	var dfs func(p int, u, x int) int

	dfs = func(p int, u, x int) int {
		// dp[u][x] = parent已经进行了x次操作2后的最大值
		if dp[u][x] > -oo {
			return dp[u][x]
		}
		// 第一种选择是，在当前节点使用操作1
		tmp1 := (coins[u] >> x) - k
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			tmp1 += dfs(u, v, x)
		}
		// 第二种选择是在当前节点使用操作2
		tmp2 := -oo
		if x+1 < H {
			tmp2 = coins[u] >> (x + 1)
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p == v {
					continue
				}
				tmp2 += dfs(u, v, x+1)
			}
		}
		dp[u][x] = max(tmp1, tmp2)

		return dp[u][x]
	}

	return dfs(-1, 0, 0)
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
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
