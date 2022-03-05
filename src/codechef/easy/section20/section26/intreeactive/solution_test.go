package main

import "testing"

const H = 20

func runSample(t *testing.T, n int, E [][]int, marked int) {
	g := buildGraph(n, E)

	D := make([]int, n)
	P := make([][]int, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			if P[u][i-1] >= 0 {
				P[u][i] = P[P[u][i-1]][i-1]
			} else {
				P[u][i] = -1
			}
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	marked--

	check := func(u, v int) bool {
		if u == marked || v == marked {
			return true
		}
		p := lca(u, v)
		if p == marked {
			return true
		}
		x := lca(u, marked)
		y := lca(v, marked)
		if x == marked {
			return lca(marked, v) == p
		}
		if y == marked {
			return lca(marked, u) == p
		}
		return false
	}

	query := func(arr []int) bool {
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr); j++ {
				u := arr[i]
				v := arr[j]
				if check(u, v) {
					return true
				}
			}
		}
		return false
	}

	res := solve(n, E, query)

	if res != marked {
		t.Errorf("Sample result %d, not correct %d", res, marked)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2},
	}
	marked := 2
	runSample(t, n, E, marked)
}

func TestSample2(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2},
	}
	marked := 1
	runSample(t, n, E, marked)
}

func TestSample3(t *testing.T) {
	n := 30
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
		{5, 15},
		{15, 16},
		{16, 17},
		{17, 18},
		{5, 11},
		{11, 12},
		{12, 13},
		{13, 14},
		{4, 21},
		{21, 20},
		{20, 19},
		{4, 22},
		{22, 23},
		{23, 24},
		{3, 27},
		{27, 28},
		{3, 25},
		{25, 26},
		{2, 29},
		{2, 30},
	}
	marked := 22
	runSample(t, n, E, marked)
}

func TestSample4(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
		{1, 6},
	}
	marked := 6
	runSample(t, n, E, marked)
}
