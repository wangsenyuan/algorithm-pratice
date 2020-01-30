package main

import "testing"

func runSample(t *testing.T, n int, a, b, c int, F []int, expect int64) {
	solver := NewSolver(a, b, c)

	for i := 0; i < n; i++ {
		solver.ReadFloor(F[i])
	}

	if solver.ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, solver.ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 5, 2, []int{6, 7, 8}, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 2, 1000000000, []int{1000000000}, 2999999997)
}
