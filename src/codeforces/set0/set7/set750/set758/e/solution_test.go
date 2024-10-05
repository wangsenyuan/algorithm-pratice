package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect bool) {
	res := solve(n, edges)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := NewGraph(n, n)

	for i, e := range res {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
	}

	var dfs func(u int) int
	dfs = func(u int) int {
		var weight int

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			e := res[g.val[i]]
			strength := e[3]
			tmp := dfs(v)
			if tmp > strength {
				t.Fatalf("Sample result %v, not correct", res)
			}
			weight += tmp + e[2]
		}
		return weight
	}

	dfs(0)
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 3, 5, 7},
		{3, 2, 4, 3},
	}
	expect := true

	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 3, 2, 3},
		{3, 4, 5, 1},
		{3, 2, 3, 3},
	}
	expect := false

	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 2, 4},
		{2, 4, 1, 9},
		{4, 5, 5, 6},
		{4, 3, 4, 8},
	}
	expect := true

	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2, 5, 2},
		{2, 3, 4, 3},
		{1, 4, 3, 7},
		{4, 5, 4, 1},
		{4, 6, 3, 2},
		{6, 7, 1, 6},
	}
	expect := true

	runSample(t, n, edges, expect)
}
