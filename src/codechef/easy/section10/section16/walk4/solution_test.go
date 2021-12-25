package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, X []int, expect []int) {
	solver := NewSolver(n, edges)

	for i := 0; i < len(X); i++ {
		res := solver.Query(X[i])
		if res != expect[i] {
			t.Fatalf("Sample %d %v, expect %v, but fail at %d with result %d", n, edges, expect, i, res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{3, 1, 3},
	}
	X := []int{1, 2, 3}
	expect := []int{2, 10, 36}
	runSample(t, n, edges, X, expect)
}
