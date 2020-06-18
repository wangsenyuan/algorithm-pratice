package main

import "testing"

func TestSample1(t *testing.T) {
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	solver := NewSolver(5, E)

	res := solver.InsertPath(4, 5)

	if res != 0 {
		t.Errorf("Sample fail at step 1, expect 0, but got %d", res)
	}

	res = solver.InsertPath(4, 2)

	if res != 1 {
		t.Errorf("Sample fail at step 2, expect 1, but got %d", res)
	}

	res = solver.InsertPath(1, 3)

	if res != 2 {
		t.Errorf("Sample fail at step 3, expect 2, but got %d", res)
	}

	res = solver.InsertPath(1, 2)

	if res != 2 {
		t.Errorf("Sample fail at step 4, expect 2, but got %d", res)
	}
}
