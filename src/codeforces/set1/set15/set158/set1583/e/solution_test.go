package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect int) {
	ans, extra := solve(n, edges, queries)

	if expect != extra {
		t.Fatalf("Sample expect %d, but got %d", expect, extra)
	}

	if expect != 0 {
		return
	}
	// expect = 0
	g := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]int)
	}

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u][v] = i
		g[v][u] = i
	}

	deg := make([]int, len(edges))

	vis := make(map[int]bool)

	for i, cur := range queries {
		a, b := cur[0], cur[1]
		path := ans[i]

		if path[0] != a || path[len(path)-1] != b {
			t.Fatalf("Sample result %v, not correct", ans)
		}

		clear(vis)

		for i := 0; i+1 < len(path); i++ {
			x := path[i] - 1
			y := path[i+1] - 1
			if j, ok := g[x][y]; !ok || vis[j] {
				t.Fatalf("Sample result %v, not correct", ans)
			} else {
				vis[j] = true
				deg[j]++
			}
		}
	}

	for i := 0; i < len(edges); i++ {
		if deg[i]%2 != 0 {
			t.Fatalf("Sample result %v, not correct", ans)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{2, 1},
		{2, 3},
		{3, 5},
		{1, 4},
		{6, 1},
		{5, 6},
		{4, 5},
	}
	queries := [][]int{
		{1, 4},
		{5, 1},
		{4, 5},
	}
	expect := 0
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{4, 3},
		{4, 5},
		{2, 1},
		{1, 4},
		{1, 3},
		{3, 5},
		{3, 2},
	}
	queries := [][]int{
		{4, 2},
		{3, 5},
		{5, 1},
		{4, 5},
	}
	expect := 2
	runSample(t, n, edges, queries, expect)
}
