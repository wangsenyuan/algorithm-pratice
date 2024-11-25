package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, r int) {
	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	lca, _ := g.prepare(r)

	var cnt int
	ask := func(u int, v int) int {
		cnt++
		return lca(u, v)
	}

	res := solve(n, edges, ask)

	if res != r {
		t.Fatalf("solution giving wrong answer %d, expecting %d", res, r)
	}
	if cnt > n/2 {
		t.Fatalf("solution took too much queries %d, exceeding %d", cnt, n/2)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 4},
		{4, 2},
		{5, 3},
		{6, 3},
		{2, 3},
	}
	expect := 4
	runSample(t, n, edges, expect)
}
func TestSample2(t *testing.T) {
	n := 9
	edges := [][]int{
		{4, 2},
		{4, 5},
		{4, 8},
		{4, 9},
		{1, 4},
		{3, 4},
		{6, 4},
		{7, 4},
	}
	expect := 4
	runSample(t, n, edges, expect)
}
