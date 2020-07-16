package main

import "testing"

func TestSample1(t *testing.T) {
	n := 5
	H := []int{3, 1, 4, 1, 5}
	S := []int{1, 2, 4, 8, 16}
	solver := NewSolver(n, H, S)

	res := solver.Query(5, 2)

	if res != 22 {
		t.Errorf("Sample query 5, 2, expect 22, but got %d", res)
	}

	solver.Update(3, 5)

	res = solver.Query(3, 4)

	if res != 13 {
		t.Errorf("Sample query 3, 4, expect 13, but got %d", res)
	}

	res = solver.Query(1, 4)

	if res != -1 {
		t.Errorf("Sample query 1, 4, expect -1, but got %d", res)
	}
}
