package main

import "testing"

func TestSample1(t *testing.T) {
	solver := NewSolver(20)

	res := solver.GetCount(1, 5)

	if res != 0 {
		t.Errorf("Sample fail at step 0")
	}

	solver.AddRangeValue(3, 5, 4)

	res = solver.GetCount(1, 5)

	if res != 1 {
		t.Errorf("Sample fail at step 1")
	}

	solver.AddRangeValue(1, 4, 15)

	res = solver.GetCount(1, 5)

	if res != 2 {
		t.Errorf("Sample fail at step 2")
	}

}
