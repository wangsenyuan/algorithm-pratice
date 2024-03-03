package p3071

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
		g.AddEdge(e[1], e[0], e[2])
	}

	var dfs func(p int, u int, dist int) int
	dfs = func(p int, u int, dist int) int {
		var cnt int
		if dist%signalSpeed == 0 {
			cnt++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt += dfs(u, v, dist+g.val[i])
			}
		}
		return cnt
	}

	ans := make([]int, n)

	for u := 0; u < n; u++ {
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			tmp := dfs(u, v, g.val[i])
			ans[u] += cnt * tmp
			cnt += tmp
		}
	}

	return ans
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
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
