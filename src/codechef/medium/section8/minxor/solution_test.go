package main

import "testing"

func TestSample1(t *testing.T) {
	arr := []int{0, 1, 0, 1, 0}
	solver := NewSolver(arr)

	res := solver.Query(1, 5)

	if res.num != 0 || res.cnt != 3 {
		t.Errorf("Fail at step 1, expect (0, 3), but got %v", res)
	}

	solver.Modify(1, 5, 1)

	res = solver.Query(1, 5)

	if res.num != 0 || res.cnt != 2 {
		t.Errorf("Fail at step 1, expect (0, 2), but got %v", res)
	}
}
