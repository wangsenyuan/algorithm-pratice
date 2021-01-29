package main

import "testing"

func runSample(t *testing.T, n int, X []uint64, E [][]int, Q []uint64, expect []uint64) {
	solver := NewSolver(n, X, E)

	for i, d := range Q {
		res := solver.Query(d)

		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
	}
	X := []uint64{1, 5, 8, 7}
	Q := []uint64{1, 2, 3}
	expect := []uint64{11, 9, 3}
	runSample(t, n, X, E, Q, expect)
}
