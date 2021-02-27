package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
	solver := NewSolver(n, E)

	for i := 0; i < len(Q); i++ {
		cur := Q[i]
		a, b := cur[0], cur[1]
		res := solver.TurnOn(a, b)
		if res != expect[i] {
			t.Fatalf("Sample expect %v, but got %d at %d", expect, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{4, 6},
	}
	Q := [][]int{
		{3, 6},
		{2, 5},
		{1, 6},
	}
	expect := []int{3, 3, 4}
	runSample(t, n, E, Q, expect)
}
