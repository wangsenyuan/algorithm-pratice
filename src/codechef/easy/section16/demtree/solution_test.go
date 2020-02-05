package main

import "testing"

func TestSample1(t *testing.T) {
	n, s := 7, 1
	R := []int{1, 1, 2, 2, 3, 3}
	W := []int{3, 5, 4, 2, 7, 9, 1}
	S := []int{1}

	solver := NewSolver(n, s, R, W, S)

	if solver.Query(2, 3, 100) != 6 {
		t.Errorf("Sample fail at 2 3 100")
	}

	if solver.Query(1, 1, 100) != 6 {
		t.Errorf("Sample fail at 1 1 100")
	}

	if solver.Query(2, 1, 100) != 6 {
		t.Errorf("Sample fail at 2 1 100")
	}

	if solver.Query(4, 5, 100) != 20 {
		t.Errorf("Sample fail at 4 5 100")
	}

	if solver.Query(4, 7, 100) != 16 {
		t.Errorf("Sample fail at 4 7 100")
	}
}
