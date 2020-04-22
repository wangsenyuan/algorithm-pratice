package main

import "testing"

func TestSample1(t *testing.T) {
	A := []int{-1, 2, -2, 1, -3}
	solver := NewSolver(len(A), A)

	res := solver.Query(3, 5)
	if res != -2 {
		t.Errorf("Sample fail at step 1, expect -2, but got %d", res)
	}

	res = solver.Query(2, 4)

	if res != 1 {
		t.Errorf("Sample fail at step 2, expect 1, but got %d", res)
	}

	solver.Update(1, 1)

	res = solver.Query(2, 4)

	if res != 2 {
		t.Errorf("Sample fail at step 3, expect 2, but got %d", res)
	}

	res = solver.Query(3, 5)

	if res != -1 {
		t.Errorf("Sample fail at step 4, expect -1, but got %d", res)
	}
}
