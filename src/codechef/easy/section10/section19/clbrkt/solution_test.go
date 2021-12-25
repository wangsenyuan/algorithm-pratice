package main

import "testing"

func TestSample1(t *testing.T) {
	s := ")())((()"
	solver := NewSolver(s)

	res := solver.Query(1)

	if res != 3 {
		t.Errorf("Sample expect 3, but got %d", res)
	}

	res = solver.Query(7)

	if res != 8 {
		t.Errorf("Sample expect 8, but got %d", res)
	}

	res = solver.Query(6)

	if res != -1 {
		t.Errorf("Sample expect -1, but got %d", res)
	}
}
