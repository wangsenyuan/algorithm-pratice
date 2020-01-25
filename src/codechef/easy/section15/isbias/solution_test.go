package main

import "testing"

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 3, 2, 4}
	solver := NewSolver(n, A)

	if solver.solve(1, 4) {
		t.Error("Sample fail at 1, 4")
	}

	if solver.solve(2, 3) {
		t.Error("Sample fail at 2, 3")
	}

	if !solver.solve(2, 4) {
		t.Error("Fail at step 2, 4")
	}
}
