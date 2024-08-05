package p3244

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	dp := make([]int, n)

	g := NewGraph(n, n*n)

	for i := 1; i < n; i++ {
		dp[i] = i
		g.AddEdge(i-1, i)
	}

	que := make([]int, n)

	bfs := func(x int, d int) {
		if dp[x] <= d {
			return
		}
		dp[x] = d

		var front, tail int
		que[front] = x
		front++

		for tail < front {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dp[v] > dp[u]+1 {
					dp[v] = dp[u] + 1
					que[front] = v
					front++
				}
			}
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		g.AddEdge(l, r)

		bfs(r, dp[l]+1)
		ans[i] = dp[n-1]
	}
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
