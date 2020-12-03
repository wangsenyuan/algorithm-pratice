package main

import "testing"

func runSample(t *testing.T, n int, P [][]int, Q [][]int, expect []int) {
	solver := NewSolver(n, P)

	for i, q := range Q {
		res := solver.Query(q[0], q[1])

		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	P := [][]int{
		{0, 0},
		{6, 0},
		{3, 4},
		{2, 1},
		{4, 1},
		{3, 3},
	}
	Q := [][]int{
		{6, 6}, {3, 3},
	}
	expect := []int{0, 1}
	runSample(t, n, P, Q, expect)
}
