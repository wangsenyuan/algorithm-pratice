package main

import "testing"

func TestSample1(t *testing.T) {
	n, m, k := 4, 5, 5
	obstacles := [][]int{
		{2, 1},
		{3, 2},
		{1, 3},
		{2, 4},
		{1, 4},
	}

	solver := NewSolver(n, m, k, obstacles)

	queries := [][]int{
		{1, 1},
		{3, 4},
		{2, 3},
		{1, 4},
	}
	expect := []int{2, 6, 3, 6}

	for i := 0; i < len(queries); i++ {
		res := solver.Query(queries[i][0], queries[i][1])

		if res != expect[i] {
			t.Fatalf("Sample fail at step %d, expect %d, but got %d", i, expect[i], res)
		}
	}
}
