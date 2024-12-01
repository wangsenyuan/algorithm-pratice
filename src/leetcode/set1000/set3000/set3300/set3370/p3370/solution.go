package p3370

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	g1 := buildGraph(edges1)
	g2 := buildGraph(edges2)

	process := func(n int, g *Graph) (cnt []int, dep []int) {
		cnt = make([]int, 2)
		dep = make([]int, n)
		var dfs func(p int, u int, d int)

		dfs = func(p int, u int, d int) {
			dep[u] = d
			cnt[d]++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					dfs(u, v, d^1)
				}
			}
		}

		dfs(-1, 0, 0)
		return
	}

	cnt1, dep1 := process(len(edges1)+1, g1)
	cnt2, _ := process(len(edges2)+1, g2)

	n := len(edges1) + 1
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = cnt1[dep1[i]]
		ans[i] += max(cnt2[0], cnt2[1])
		// 假设i连接到节点j，那么增加的答案 = 距离这个节点为偶数距离
	}
	return ans
}

func buildGraph(edges [][]int) *Graph {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}
	return g
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
