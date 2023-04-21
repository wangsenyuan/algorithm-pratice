package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	root, P := solve(n, E)
	root--

	g := createGraphFromEdges(n, E)

	var dfs func(p int, u int)

	big := make([]int, n)
	D := make([]int, n)
	sz := make([]int, n)
	fa := make([]int, n)
	dfs = func(p int, u int) {
		big[u] = -1
		sz[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				fa[v] = u
				dfs(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[v] > sz[big[u]] {
					big[u] = v
				}
			}
		}
	}

	dfs(-1, root)

	bl := make([]int, n)

	var dfs2 func(p int, u int, x int)

	dfs2 = func(p int, u int, x int) {
		bl[u] = x
		if big[u] >= 0 {
			dfs2(u, big[u], x)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && big[u] != v {
				dfs2(u, v, v)
			}
		}
	}

	dfs2(-1, root, root)

	lca := func(u, v int) int {
		for bl[u] != bl[v] {
			if D[bl[u]] < D[bl[v]] {
				u, v = v, u
			}
			u = fa[bl[u]]
		}
		if D[u] > D[v] {
			u, v = v, u
		}
		// D[u] < D[v]
		return u
	}

	set := make(map[int]bool)

	for i := 0; i+1 < n; i++ {
		u, v := P[i], P[i+1]
		p := lca(u-1, v-1)
		set[p] = true
	}

	if len(set) != expect {
		t.Fatalf("Sample expect %d sized-set, but got %v with root %d, P %v", expect, set, root+1, P)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{{1, 2}}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 1
	runSample(t, n, E, expect)
}
