package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int) {
	res := solve(k, n, edges)

	if len(res) <= k {
		t.Fatalf("Sample result %v, not correct", res)
	}

	g := createGraph(n, edges)

	u := res[0]

	for i := 1; i < len(res); i++ {
		v := res[i]
		if g[u][v] == 0 {
			t.Fatalf("Sample result %v, has no road between %d %d", res, u, v)
		}
		u = v
	}

	if g[u][res[0]] == 0 {
		t.Fatalf("Sample result %v, not returning back.", res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	runSample(t, n, edges, k)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	edges := [][]int{
		{4, 3},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
	}
	runSample(t, n, edges, k)
}

func TestSample3(t *testing.T) {
	n, k := 10, 3
	edges := [][]int{
		{9, 4},
		{4, 8},
		{4, 2},
		{2, 9},
		{9, 6},
		{6, 2},
		{6, 5},
		{1, 10},
		{1, 7},
		{10, 5},
		{10, 7},
		{1, 8},
		{8, 3},
		{3, 5},
		{3, 7},
	}
	runSample(t, n, edges, k)
}
