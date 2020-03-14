package main

import "testing"

func TestSample1(t *testing.T) {
	n := 5
	A := []int64{1, 2, 3, 4, 5}

	solver := NewSolver(n, A)

	for i := (n - 1) / BL; i >= 0; i-- {
		solver.UpdateBlock(i)
	}

	if solver.Query(0, 2) != 2 {
		t.Errorf("Fail at step 1")
	}

	solver.EditRange(2, 3, -1)

	if solver.Query(0, 2) != 3 {
		t.Errorf("Fail at step 2")
	}
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int64{1, 2, 3, 4, 5}

	solver := NewSolver1(n, A)

	if solver.Query(0, 2) != 2 {
		t.Errorf("Fail at step 1")
	}

	solver.EditRange(2, 3, -1)

	if solver.Query(0, 2) != 3 {
		t.Errorf("Fail at step 2")
	}
}
