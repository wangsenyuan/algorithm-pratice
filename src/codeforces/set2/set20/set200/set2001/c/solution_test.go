package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int) {
	var cnt int

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	fa := make([][]int, n+1)
	fa[0] = make([]int, 10)
	dep := make([]int, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, 10)
		fa[u][0] = p
		for i := 1; i < 10; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 1)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := 9; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				u = fa[u][i]
			}
		}

		if u == v {
			return u
		}
		for i := 9; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	kth := func(u int, k int) int {
		for i := 9; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = fa[u][i]
			}
		}
		return u
	}

	ask := func(a, b int) int {
		cnt++
		if a == b {
			return a
		}

		p := lca(a, b)
		dis := dep[a] + dep[b] - 2*dep[p]

		u := a
		if dep[a] < dep[b] {
			u = b
		}

		v := kth(u, dis/2)

		if dis%2 == 1 && u == b {
			return fa[v][0]
		}
		return v
	}

	res := solve(n, ask)

	if cnt > 15*n {
		t.Fatalf("Sammple result asked too much %d", cnt)
	}

	if len(res) != n-1 {
		t.Fatalf("Sample result %v, not a tree", res)
	}

	tr := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		tr[i] = make(map[int]bool)
	}

	for _, cur := range res {
		u, v := cur[0], cur[1]
		tr[u][v] = true
		tr[v][u] = true
	}

	var verify func(p int, u int) bool

	verify = func(p int, u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if !tr[u][v] || !verify(u, v) {
				return false
			}
		}
		return true
	}

	if !verify(0, 1) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2}, {1, 3}, {3, 4},
	}
	runSample(t, n, edges)
}

func TestSample2(t *testing.T) {
	n := 1000

	var edges [][]int

	nodes := []int{1}

	for i := 2; i <= n; i++ {
		j := rand.Intn(len(nodes))
		edges = append(edges, []int{nodes[j], i})
		nodes = append(nodes, i)
	}

	runSample(t, n, edges)
}
