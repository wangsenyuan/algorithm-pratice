package p2603

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)

	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs1 func(p int, u int)

	mark := make([]bool, n)
	deg := make([]int, n)

	cnt := make([]int, n)

	dfs1 = func(p int, u int) {
		if p >= 0 {
			deg[u]++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				if cnt[v] == 0 {
					mark[v] = true
				} else {
					deg[u]++
				}
				cnt[u] += cnt[v]
			}
		}
		cnt[u] += coins[u]
	}

	r := -1
	for i := 0; i < n; i++ {
		if coins[i] == 1 {
			r = i
			break
		}
	}
	if r < 0 {
		// no coins
		return 0
	}
	dfs1(-1, r)

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

	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = -1
		if coins[i] == 1 && deg[i] == 1 {
			push(i)
			L[i] = 0
		}
	}
	if front == end {
		return 0
	}

	var res int
	for front < end {
		u := pop()
		if L[u] >= 2 {
			res++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if !mark[v] && deg[v] <= 1 && L[v] < 0 {
				L[v] = L[u] + 1
				push(v)
			}
		}
	}
	if res == 0 {
		return 0
	}
	return 2 * (res - 1)
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
