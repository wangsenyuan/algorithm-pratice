package p2050

func minimumTime(n int, relations [][]int, time []int) int {

	g := NewGraph(n+3, len(relations)+2*n+2)
	r := NewGraph(n+3, len(relations)+2*n+2)
	for _, relation := range relations {
		u, v := relation[0], relation[1]
		g.AddEdge(v, u)
		r.AddEdge(u, v)
	}

	for i := 1; i <= n; i++ {
		r.AddEdge(i, n+1)
		r.AddEdge(0, i)
		g.AddEdge(n+1, i)
		g.AddEdge(i, 0)
	}

	arr := make([]int, 0, n+2)
	vis := make([]bool, n+2)
	var topo func(u int)

	topo = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				topo(v)
			}
		}
		arr = append(arr, u)
	}

	topo(n + 1)

	res := make([]int, n+2)

	for _, u := range arr {
		for j := r.nodes[u]; j > 0; j = r.next[j] {
			v := r.to[j]
			if v <= n {
				res[v] = max(res[v], res[u]+time[v-1])
			} else {
				res[v] = max(res[v], res[u])
			}
		}
	}
	return res[n+1]
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
