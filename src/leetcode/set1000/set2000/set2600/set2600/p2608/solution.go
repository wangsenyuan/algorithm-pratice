package p2608

func findShortestCycle(n int, edges [][]int) int {
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
		g.AddEdge(e[1], e[0])
	}

	que := make([]int, n)
	var front, end int

	push := func(u int) {
		que[end] = u
		end++
	}

	pop := func() int {
		u := que[front]
		front++
		return u
	}

	dist := make([]int, n)
	prev := make([]int, n)
	bfs := func(x int) int {
		for i := 0; i < n; i++ {
			dist[i] = -1
			prev[i] = -1
		}
		tmp := -1
		front = 0
		end = 0
		dist[x] = 0
		push(x)
		for front < end {
			u := pop()
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]

				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					prev[v] = u
					push(v)
				} else if prev[u] != v && dist[u]+dist[v]+1 >= 3 {
					// dist[v] > 0
					if tmp < 0 || tmp > dist[u]+dist[v]+1 {
						tmp = dist[u] + dist[v] + 1
					}
				}
			}
		}

		return tmp
	}

	best := -1

	for i := 0; i < n; i++ {
		tmp := bfs(i)
		if tmp > 0 && (best < 0 || best > tmp) {
			best = tmp
		}
	}

	return best
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
