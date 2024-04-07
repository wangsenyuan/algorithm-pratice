package p3105

func minimumCost(n int, edges [][]int, query [][]int) []int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	belong := make([]int, n)
	for i := 0; i < n; i++ {
		belong[i] = -1
	}
	var comp []int
	var dfs func(u int) int
	dfs = func(u int) int {
		res := (1 << 30) - 1
		if belong[u] >= 0 {
			return res
		}

		belong[u] = len(comp)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			res = res & g.val[i] & dfs(v)
		}
		return res
	}

	for i := 0; i < n; i++ {
		if belong[i] < 0 {
			comp = append(comp, dfs(i))
		}
	}

	ans := make([]int, len(query))

	for i, cur := range query {
		u, v := cur[0], cur[1]
		if u == v {
			ans[i] = 0
		} else {
			if belong[u] != belong[v] {
				ans[i] = -1
			} else {
				ans[i] = comp[belong[u]]
			}
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
	e++
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
