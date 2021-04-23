package main

import "testing"

func TestSample1(t *testing.T) {
	A := []int{0, 0, 0, 0, 0}
	B := []int{1, 2, 3, 4, 5}
	solver := NewSolver(A, B)
	solver.Update(2, 3, 1)
	res := solver.Get(1, 5)
	if res != 3 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}
	solver.UpdatePers(3, 2)

	res = solver.Get(1, 5)
	if res != 3 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}
	solver.Update(1, 5, 1)
	res = solver.Get(1, 5)
	if res != 14 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}
}
