package p2468

import "math"

const H = 20

func mostProfitablePath(edges [][]int, bob int, amount []int) int {

	n := len(edges) + 1

	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	leaf := make([]bool, n)
	amt := make([]int, n)
	D := make([]int, n)
	anc := make([][]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		anc[u] = make([]int, H)
		anc[u][0] = p
		for i := 1; i < H; i++ {
			anc[u][i] = anc[anc[u][i-1]][i-1]
		}
		leaf[u] = true
		amt[u] += amount[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				leaf[u] = false
				amt[v] = amt[u]
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = anc[u][i]
			}
		}

		if u == v {
			return u
		}

		for i := H - 1; i >= 0; i-- {
			if anc[u][i] != anc[v][i] {
				u = anc[u][i]
				v = anc[v][i]
			}
		}
		return anc[u][0]
	}

	kth := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if k >= 1<<i {
				u = anc[u][i]
				k -= 1 << i
			}
		}
		return u
	}

	mid := kth(bob, (D[bob]+1)/2)
	mid_cost := amt[mid]

	if D[bob]%2 == 0 {
		// meet in the mid
		mid_cost -= amount[mid] / 2
	}

	var best int = math.MinInt32

	for i := 0; i < n; i++ {
		if leaf[i] {
			p := lca(i, bob)
			x := lca(p, mid)
			if x == mid {
				best = max(best, amt[i]-amt[p]+mid_cost)
			} else {
				best = max(best, amt[i]) 
			}
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
