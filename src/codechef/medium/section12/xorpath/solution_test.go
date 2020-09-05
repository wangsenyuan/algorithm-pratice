package main

import "testing"

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2, 36},
		{2, 4, 16},
		{4, 5, 18},
		{2, 3, 50},
		{3, 4, 11},
		{6, 6, 100},
		{7, 6, 100},
	}
	solver := NewSolver(n, E)

	queries := [][]int{
		{1, 5},
		{5, 1},
		{1, 2},
		{6, 7},
		{1, 7},
	}
	expect := []int64{15, 15, 13, 0, -1}

	for i := 0; i < len(queries); i++ {
		ans := solver.Ask(queries[i][0], queries[i][1])
		if ans != expect[i] {
			t.Errorf("Sample expect %d, but got %d", expect[i], ans)
		}
	}
}
