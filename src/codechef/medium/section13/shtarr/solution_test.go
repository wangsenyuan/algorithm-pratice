package main

import "testing"

func TestSample1(t *testing.T) {
	n := 6
	A := []int{3, 7, 2, 4, 8, 7}
	solver := NewSolver(n, A)

	res := solver.GetAnswer(4, 3, 9)
	expect(t, res == 2, "result should be 2")
	res = solver.GetAnswer(3, 1, 7)
	expect(t, res == 3, "result should be 3")
	solver.Add(6, 2)
	res = solver.GetAnswer(4, 3, 9)
	expect(t, res == 3, "result should be 3")
}

func expect(t *testing.T, expr bool, message string) {
	if !expr {
		t.Fatalf("Sample fail %s", message)
	}
}
