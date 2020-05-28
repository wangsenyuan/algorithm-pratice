package main

import "testing"

func TestSample1(t *testing.T) {
	N := 10
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
		{3, 10},
		{7, 8},
		{7, 9},
	}
	C := []int{1, 2, 2, 1, 1, 3, 4, 4, 3, 4}
	A := []int{1, 2, 2, 1, 1, 3, 4, 4, 3, 4}

	solver := NewSolver(N, E, C, A)

	ops := [][]int{
		{1, 1},
		{2, 3},
		{4, 2},
		{2, 1},
		{8, 2},
		{10, 3},
	}
	expect := []int64{38, 38, 40, 38, 42, 42}

	for i := 0; i < len(ops); i++ {
		op := ops[i]
		res := solver.ChangeColor(op[0], op[1])

		if res != expect[i] {
			t.Errorf("Sample fail at %d, expect %d, but got %d", i, expect[i], res)
		}
	}
}
