package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect bool) {
	res := solve(n, E)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	ok := make([]bool, n)
	vis := make(map[int]bool)
	for _, u := range res {
		u--
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if ok[v] {
				t.Fatalf("Sample result %v at %d - %d, violates condition 3", res, u+1, v+1)
			}
			vis[v] = true
		}
		ok[u] = true
	}

	if len(vis) != n {
		t.Fatalf("Sample result %v, violates conditions", res)
	}
}

func TestSample1(t *testing.T) {
	n := 17
	E := [][]int{
		{1, 8},
		{2, 9},
		{3, 10},
		{4, 11},
		{5, 12},
		{6, 13},
		{7, 14},
		{8, 9},
		{8, 14},
		{8, 15},
		{9, 10},
		{9, 15},
		{10, 11},
		{10, 15},
		{10, 17},
		{11, 12},
		{11, 17},
		{12, 13},
		{12, 16},
		{12, 17},
		{13, 14},
		{13, 16},
		{14, 16},
		{14, 15},
		{15, 16},
		{15, 17},
		{16, 17},
	}
	expect := true
	runSample(t, n, E, expect)
}
