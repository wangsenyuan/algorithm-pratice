package p2049

func countHighestScoreNodes(parents []int) int {
	n := len(parents)

	g := NewGraph(n, n+2)

	for i := 1; i < n; i++ {
		g.AddEdge(parents[i], i)
	}
	sz := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs(g.to[i])
			sz[u] += sz[g.to[i]]
		}
	}

	dfs(0)

	score := make([]int64, n)

	var dfs2 func(u int)

	dfs2 = func(u int) {
		if u > 0 {
			score[u] = int64(n - sz[u])
		} else {
			score[u] = 1
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs2(v)
			score[u] *= int64(sz[v])
		}
	}

	dfs2(0)

	var ans int
	var cnt int = 1

	for i := 1; i < n; i++ {
		if score[i] > score[ans] {
			ans = i
			cnt = 1
		} else if score[i] == score[ans] {
			cnt++
		}
	}

	return cnt
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
