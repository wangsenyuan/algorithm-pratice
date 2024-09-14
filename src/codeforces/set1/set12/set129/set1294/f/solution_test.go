package main

import (
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	cnt, res := solve(n, edges)

	y := check(n, edges, res)

	if cnt != y || cnt != expect {
		t.Fatalf("Sample expect %v, but got %d, %v => %d", expect, cnt, res, y)
	}
}

func check(n int, edges [][]int, res []int) int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	a, b, c := res[0]-1, res[1]-1, res[2]-1

	var cnt int

	var dfs func(p int, u int) bool

	dfs = func(p int, u int) bool {
		if u == b || u == c {
			return true
		}
		ok := false
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if dfs(u, v) {
					ok = true
					cnt++
				}
			}
		}
		return ok
	}

	dfs(-1, a)

	return cnt
}

func TestSample1(t *testing.T) {
	n := 8
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 6},
		{3, 7},
		{3, 8},
	}
	expect := 5
	runSample(t, n, edges, expect)
}
