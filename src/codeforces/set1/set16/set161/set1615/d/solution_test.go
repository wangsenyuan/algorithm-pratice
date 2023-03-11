package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, X [][]int, valid bool) {
	res := solve(n, E, X)

	if valid != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", valid, res)
	}

	if !valid {
		return
	}

	g := createGraphFromEdges(n, res)

	var dfs func(p int, u int)

	W := make([]int, n)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				W[v] = W[u] ^ g.val[i]
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	for _, e := range X {
		u, v, w := e[0]-1, e[1]-1, e[2]
		x := W[u] ^ W[v]
		if digit_count_xor(x) != w {
			t.Fatalf("Sample result not correct")
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2, -1},
		{1, 3, 1},
		{4, 2, 7},
		{6, 3, 0},
		{2, 5, -1},
	}
	X := [][]int{
		{2, 3, 1},
		{2, 5, 0},
		{5, 6, 1},
		{6, 1, 1},
		{4, 5, 1},
	}
	valid := true
	runSample(t, n, E, X, valid)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, -1},
		{1, 3, -1},
		{1, 4, 1},
		{4, 5, -1},
	}
	X := [][]int{
		{2, 4, 0},
		{3, 4, 1},
		{2, 3, 1},
	}
	valid := true
	runSample(t, n, E, X, valid)
}

func TestSample3(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, -1},
		{1, 3, -1},
	}
	X := [][]int{
		{1, 2, 0},
		{1, 3, 1},
		{2, 3, 0},
	}
	valid := false
	runSample(t, n, E, X, valid)
}

func TestSample4(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 1},
	}
	X := [][]int{
		{1, 2, 0},
	}
	valid := false
	runSample(t, n, E, X, valid)
}
