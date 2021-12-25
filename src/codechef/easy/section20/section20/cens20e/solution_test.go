package main

import "testing"

func runSample(t *testing.T, n int, m int, W []int, E [][]int, queries [][]int, expect []int) {
	solver := NewSolver(n, m, W, E)

	for i, query := range queries {
		a, k := query[0], query[1]
		ans := solver.Query(a, k)
		if ans != expect[i] {
			t.Errorf("Sample fail at %d, expect %d, but got %d", i, expect[i], ans)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 6
	W := []int{6, 4, 10, 7, 5, 5}
	E := [][]int{
		{3, 4},
		{5, 1},
		{5, 6},
		{3, 6},
		{1, 4},
		{5, 3},
	}
	queries := [][]int{
		{1, 6},
		{1, 8},
		{1, 7},
		{2, 6},
	}
	expect := []int{1, 4, 4, 4}
	runSample(t, n, m, W, E, queries, expect)
}
