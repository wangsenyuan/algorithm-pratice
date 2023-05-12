package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, Q []int, expect [][]int) {
	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	res := solve(n, E, Q)

	P := make([]Pair, n)
	D := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				P[v] = Pair{u, g.val[i]}
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	for i := 0; i < len(res); i++ {
		a := res[i]
		b := expect[i]
		if checkResult(D, P, a[0], a[1]) != checkResult(D, P, b[0], b[1]) {
			t.Fatalf("Sample result %v, not correct, expect %v", a, b)
		}
	}
}

func checkResult(D []int, P []Pair, u int, v int) int {
	if u < 0 {
		return 0
	}
	var cnt int
	var and int
	for u != v {
		and &= P[u].second
		and &= P[v].second
		if and == 0 {
			return -1
		}
		if D[u] > D[v] {
			u = P[u].first
		} else {
			v = P[v].first
		}
	}

	return cnt + 1
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{2, 1, 5},
		{3, 2, 2},
		{4, 2, 4},
		{5, 4, 6},
		{6, 4, 0},
	}
	Q := []int{2, 3, 6}
	expect := [][]int{
		{5, 1},
		{3, 2},
		{-1, -1},
	}
	runSample(t, n, E, Q, expect)
}
