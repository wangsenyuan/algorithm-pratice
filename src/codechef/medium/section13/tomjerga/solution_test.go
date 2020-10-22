package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, queries [][]int, expect []int) {
	solver := NewSolver(n, E)

	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		ans := solver.Solve(y, x)
		if ans != expect[i] {
			t.Fatalf("Sample fail at step %d, expect %d, but got %d", i, expect[i], ans)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	queries := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := []int{2, 1}
	runSample(t, n, E, queries, expect)
}
