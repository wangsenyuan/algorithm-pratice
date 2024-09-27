package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect int) {
	cnt, res := solve(n, roads)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}

	if expect <= 0 {
		return
	}

	g := NewGraph(n, len(roads)*2+2*cnt)

	for _, e := range roads {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		if u == v {
			t.Fatalf("Sample result %v, not correct, having self loop", res)
		}
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	low := make([]int, n)
	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = -1
	}

	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = timer
		low[u] = timer
		timer++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if in[v] < 0 {
				dfs(u, v)
				if low[v] == in[v] {
					// a bridge
					t.Fatalf("Sample result %v, not correct, as the graph still has bridge", res)
				}
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], in[v])
			}
		}
	}

	dfs(-1, 0)
}

func TestSample1(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 1
	runSample(t, n, roads, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	roads := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := 1
	runSample(t, n, roads, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	roads := [][]int{
		{6, 4},
		{3, 7},
		{4, 9},
		{8, 4},
		{3, 4},
		{3, 6},
		{7, 5},
		{3, 9},
		{10, 9},
		{10, 5},
		{1, 2},
		{1, 8},
		{8, 2},
		{5, 6},
		{6, 9},
		{5, 9},
		{3, 10},
		{7, 10},
	}
	expect := 1
	runSample(t, n, roads, expect)
}
