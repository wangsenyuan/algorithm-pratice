package main

import "testing"

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
	}

	solver := NewSolver(n, edges)

	res := solver.Query(2, 1, 4, 1)

	if res != 3 {
		t.Errorf("Sample fail at step 1 => %d", res)
	}

	res = solver.Query(2, 2, 4, 2)

	if res != 5 {
		t.Errorf("Sample fail at step 2 => %d", res)
	}

	res = solver.Query(1, 1, 2, 1)

	if res != -1 {
		t.Errorf("Sample fail at step 3 => %d", res)
	}
}
