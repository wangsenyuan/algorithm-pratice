package main

import "testing"

func runSample(t *testing.T, n int, c [][]int, expect []int) {
	res := solve(n, c)

	x := check(n, c, expect)
	y := check(n, c, res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func check(n int, c [][]int, expect []int) int {
	g := NewGraph(n, 2*n)

	for i := 0; i < n; i++ {
		p := expect[i]
		if p > 0 {
			g.AddEdge(p-1, i)
			g.AddEdge(i, p-1)
		}
	}

	var res int

	var dfs func(p int, u int, d int, from int)
	dfs = func(p int, u int, d int, from int) {
		res += c[from][u] * d
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, d+1, from)
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(-1, i, 0, i)
	}

	return res
}

func TestSample1(t *testing.T) {
	n := 4
	c := [][]int{
		{0, 566, 1, 0},
		{566, 0, 239, 30},
		{1, 239, 0, 1},
		{0, 30, 1, 0},
	}
	expect := []int{2, 4, 2, 0}
	runSample(t, n, c, expect)
}
