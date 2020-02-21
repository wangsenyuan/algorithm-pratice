package main

import "testing"

func TestSample1(t *testing.T) {
	s := "AA"
	solver := NewSolver(len(s), s)

	solver.UpdateRange(2, 2, 'B')

	p, q := solver.Query(1, 1, 1, 1)
	if p != 0 || q != 2 {
		t.Errorf("Sample fail at step 0")
	}

	p, q = solver.Query(1, 2, 1, 2)

	if p != 2 || q != 4 {
		t.Errorf("Sample fail at step 1")
	}
}
