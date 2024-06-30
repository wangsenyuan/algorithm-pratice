package p3202

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	da, a := getDiameter(edges1)
	db, b := getDiameter(edges2)
	best := 1 << 30

	for _, x := range a {
		for _, y := range b {
			best = min(best, x.second+y.second)
		}
	}

	return max(best, da, db) - 1
}

type Pair struct {
	first  int
	second int
}

func getDiameter(edges [][]int) (int, []Pair) {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	dist := make([]int, n)
	track := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		track[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] == 0 {
				dist[v] = dist[u] + 1
				dfs(u, v)
			}
		}
	}
	dist[0] = 1
	dfs(-1, 0)

	var far int
	for u := 0; u < n; u++ {
		if dist[u] > dist[far] {
			far = u
		}
	}
	clear(dist)

	dist[far] = 1

	dfs(-1, far)

	second := far
	for u := 0; u < n; u++ {
		if dist[u] > dist[second] {
			second = u
		}
	}

	dia := dist[second]

	h := dia/2 + 1

	for dist[second] != h {
		second = track[second]
	}
	res := []Pair{Pair{second, h}}
	if dia%2 == 0 {
		oth := track[second]
		res = append(res, Pair{oth, dia + 1 - dist[oth]})
	}
	return dia, res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
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
