package p2538

// const INF = 1 << 50

func maxOutput(n int, edges [][]int, price []int) int64 {
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	fp := make([]int64, n)

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				fp[u] = max(fp[u], fp[v]+int64(price[u]))
			}
		}
		if fp[u] == 0 {
			fp[u] = int64(price[u])
		}
	}

	dfs1(0, 0)

	var res int64

	// dfs2 计算经过u的最大值
	var dfs2 func(p int, u int, sum int64)
	dfs2 = func(p int, u int, sum int64) {
		a, b := -1, -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				if a < 0 || fp[a] < fp[v] {
					b = a
					a = v
				} else if b < 0 || fp[b] < fp[v] {
					b = v
				}
			}
		}
		if a < 0 {
			// leaf
			res = max(res, sum)
			return
		}
		res = max(res, max(sum, fp[a]))

		if b < 0 {
			// only one child
			dfs2(u, a, sum+int64(price[u]))
			return
		}
		// more children
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if v == a {
					dfs2(u, v, max(sum, fp[b])+int64(price[u]))
				} else {
					dfs2(u, v, max(sum, fp[a])+int64(price[u]))
				}
			}
		}
	}

	dfs2(0, 0, 0)

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
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
