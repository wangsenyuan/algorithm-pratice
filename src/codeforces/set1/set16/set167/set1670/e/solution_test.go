package main

import "testing"

func runSample(t *testing.T, p int, E [][]int, root int, V []int, X []int) {
	r, v, x := solve(p, E)
	n := len(E) + 1
	g := NewGraph(n, 2*n)
	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	expect := check(root-1, V, X, g)
	ans := check(r-1, v, x, g)

	if expect != ans {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func check(root int, V []int, X []int, g *Graph) int {
	var ans int

	var dfs func(p int, u int, cur int)

	dfs = func(p int, u int, cur int) {
		ans = max(ans, cur)
		cur ^= V[u]
		ans = max(ans, cur)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if p != v {
				dfs(u, v, cur^X[w])
			}
		}
	}

	dfs(root, root, 0)

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	p := 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	V := []int{5, 1, 3, 6}
	X := []int{4, 2, 7}
	root := 3
	runSample(t, p, E, root, V, X)
}

func TestSample2(t *testing.T) {
	p := 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 5},
		{1, 6},
		{5, 7},
		{5, 8},
	}
	V := []int{1, 2, 8, 11, 4, 13, 9, 15}
	X := []int{6, 14, 3, 7, 10, 5, 12}
	root := 5
	runSample(t, p, E, root, V, X)
}
