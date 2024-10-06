package p3310

func remainingMethods(n int, k int, invocations [][]int) []int {
	// 先从k出发，把所有它能访问到的都比较出来
	// 然后从任何未标记的方法出发，再标记一遍
	g := NewGraph(n, len(invocations)+1)

	for _, cur := range invocations {
		u, v := cur[0], cur[1]
		g.AddEdge(u, v)
	}

	bug := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		bug[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !bug[v] {
				dfs(v)
			}
		}
	}
	dfs(k)

	marked := make([]bool, n)

	var dfs2 func(u int) bool
	dfs2 = func(u int) bool {
		if bug[u] {
			return true
		}
		marked[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !marked[v] {
				if dfs2(v) {
					return true
				}
			}
		}
		return false
	}
	var res []int

	for i := 0; i < n; i++ {
		if !bug[i] && !marked[i] {
			if dfs2(i) {
				// 访问到了有bug的程序
				for j := 0; j < n; j++ {
					res = append(res, j)
				}
				return res
			}
		}
	}

	for i := 0; i < n; i++ {
		if marked[i] {
			res = append(res, i)
		}
	}

	return res
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
