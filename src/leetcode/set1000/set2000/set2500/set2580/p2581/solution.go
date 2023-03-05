package p2581

func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1

	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	in := make([]int, n)
	out := make([]int, n)
	L := make([]int, n)
	var time int
	var dfs1 func(p int, u int)
	dfs1 = func(p int, u int) {
		in[u] = time
		time++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				L[v] = L[u] + 1
				dfs1(u, v)
			}
		}
		out[u] = time
	}

	dfs1(0, 0)

	arr := make([]int, n*2)

	add := func(l int, r int, v int) {
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				arr[l] += v
				l++
			}
			if r&1 == 1 {
				r--
				arr[r] += v
			}
			l >>= 1
			r >>= 1
		}
	}

	get := func(p int) int {
		p += n
		var res int
		for p > 0 {
			res += arr[p]
			p >>= 1
		}
		return res
	}

	for _, gu := range guesses {
		u, v := gu[0], gu[1]
		if L[u] < L[v] {
			// u is parent v
			add(0, n, 1)
			add(in[v], out[v], -1)
		} else {
			// v is parent,
			add(in[u], out[u], 1)
		}
	}
	var res int
	for i := 0; i < n; i++ {
		x := get(i)
		if x >= k {
			res++
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
