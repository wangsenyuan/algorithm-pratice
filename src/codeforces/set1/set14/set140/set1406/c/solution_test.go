package main

import "testing"

func runSample(t *testing.T, n int, E [][]int) {
	res := solve(n, E)

	g := NewGraph(n, n*2)

	for _, e := range E {
		if e[0] == res[0][0] && e[1] == res[0][1] {
			continue
		}
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	u, v := res[1][0]-1, res[1][1]-1
	g.AddEdge(u, v)
	g.AddEdge(v, u)

	sz := make([]int, n)
	par := make([]int, n)
	x := findCentroid(n, sz, par, g)

	if sz[x]*2 == n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
	}
	runSample(t, n, E)
}
