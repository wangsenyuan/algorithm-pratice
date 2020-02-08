package main

import "testing"

func TestSample1(t *testing.T) {
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 4, 1},
		{4, 1, 1},
	}
	solver := NewSolver(4, edges)

	if solver.GetAnswer() != 3 {
		t.Errorf("Sample fail at first step")
	}

	solver.AssigneZero(1, 2)
	if solver.GetAnswer() != 2 {
		t.Errorf("Sample fail at second step")
	}

	solver.AssignOriginal(1, 2)

	if solver.GetAnswer() != 3 {
		t.Errorf("Sample fail at third step")
	}
}
