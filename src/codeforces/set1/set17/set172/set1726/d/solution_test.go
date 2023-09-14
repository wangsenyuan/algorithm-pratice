package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect string) {
	res := solve(n, E)

	a := check(n, E, expect, '0') + check(n, E, expect, '1')
	b := check(n, E, res, '0') + check(n, E, res, '1')

	if a != b {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func check(n int, E [][]int, color string, exp byte) int {
	g := NewGraph(n, len(E)*2)
	for i, e := range E {
		if color[i] == exp {
			u, v := e[0]-1, e[1]-1
			g.AddEdge(u, v, i)
			g.AddEdge(v, u, i)
		}
	}
	vis := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
			ans++
		}
	}
	return ans
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 1},
		{1, 3},
		{3, 5},
	}
	expect := "0111010"
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{3, 4},
	}
	expect := "1001"
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
		{1, 4},
		{5, 6},
		{6, 2},
	}
	expect := "0001111"
	runSample(t, n, E, expect)
}
