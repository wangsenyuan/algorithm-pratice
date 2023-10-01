package p2878

func countVisitedNodes(edges []int) []int {
	n := len(edges)

	g := NewGraph(n, n+1)

	for i := 0; i < n; i++ {
		j := edges[i]
		g.AddEdge(j, i)
	}

	vis := make([]int, n)
	for i := 0; i < n; i++ {
		vis[i] = -1
	}
	trace := make([]int, n)
	inCycle := make([]bool, n)
	cycles := make([]int, n)
	ans := make([]int, n)

	for u := 0; u < n; u++ {
		if vis[u] >= 0 {
			continue
		}
		cur := u
		for vis[cur] < 0 {
			vis[cur] = u
			trace[edges[cur]] = cur
			cur = edges[cur]
		}
		if vis[cur] != u {
			continue
		}
		var cnt int
		for !inCycle[cur] {
			inCycle[cur] = true
			cycles[cnt] = cur
			cnt++
			cur = trace[cur]
		}
		for i := 0; i < cnt; i++ {
			ans[cycles[i]] = cnt
		}
	}

	// from cur go back to trace, will be in cycle

	var dfs func(u int)

	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if inCycle[v] {
				continue
			}
			ans[v] = ans[u] + 1
			dfs(v)
		}
	}

	for i := 0; i < n; i++ {
		if inCycle[i] {
			dfs(i)
		}
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
