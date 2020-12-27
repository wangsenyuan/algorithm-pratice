package main

import "testing"

func runSample(t *testing.T, n int, S string, Q [][]int, expect []int) {
	solver := NewSolver(n, S)
	for i, q := range Q {
		res := solver.Solve(q[0], q[1], q[2])
		if q[0] == 1 {
			continue
		}
		if res != expect[i] {
			t.Fatalf("Sample result %d at step %d, not correct", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 7
	S := "abacaba"
	Q := [][]int{
		{2, 1, 7},
		{2, 2, 3},
		{1, 1, 2},
		{2, 2, 3},
	}
	expect := []int{3, 0, 0, 1}
	runSample(t, n, S, Q, expect)
}
