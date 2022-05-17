package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, X []int) {
	res := solve(n, E, X)

	g := buildTree(n, E)

	cnt := make([][]int, n)

	dfs(-1, 0, cnt, X, g)

	diff := cnt[0][1] - cnt[0][0]

	for _, v := range res {
		v--
		diff += 2 * (cnt[v][0] - cnt[v][1])
	}

	if diff != 0 {
		t.Errorf("Sample result %v, not correct, after flip, diff is %d", res, diff)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	X := []int{1, 1, 1, 1}
	runSample(t, n, E, X)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	X := []int{0, 1, 0, 0}
	runSample(t, n, E, X)
}

func TestSample3(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	X := []int{0, 1, 0, 1}
	runSample(t, n, E, X)
}
