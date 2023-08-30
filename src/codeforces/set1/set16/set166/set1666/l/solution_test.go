package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, s int, expect bool) {
	res := solve(n, E, s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	g := NewGraph(n, len(E))
	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
	}
	a := check(g, s, res[0])
	b := check(g, s, res[1])
	if a < 0 || b < 0 {
		t.Fatalf("Sample result not correct, %v", res)
	}
}

func check(g *Graph, s int, path []int) int {

	var dfs func(u int, j int) int

	dfs = func(u int, j int) int {
		if j+1 == len(path) {
			return j + 1
		}
		// path[i] = u
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v+1 == path[j+1] {
				return dfs(v, j+1)
			}
		}

		return -1
	}
	if path[0] != s {
		return -1
	}
	return dfs(s-1, 0)
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 3},
		{3, 5},
	}
	s := 1
	expect := true
	runSample(t, n, E, s, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{2, 5},
		{5, 4},
	}
	s := 1
	expect := false
	runSample(t, n, E, s, expect)
}
