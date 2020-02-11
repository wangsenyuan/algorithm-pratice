package main

import "testing"

func runSample(t *testing.T, n int, A []int, M []int, expect []int) {
	solver := NewSolver(n, A)

	for i := 0; i < len(M); i++ {
		if solver.Query(M[i]) != expect[i] {
			t.Errorf("Sample %d %v, failed at step %d, expect %d, but got %d", n, A, i, expect[i], solver.Query(M[i]))
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, -1, 4, 5}
	M := []int{9}
	expect := []int{4}
	runSample(t, n, A, M, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	M := []int{5, 15}
	expect := []int{8, 2}
	runSample(t, n, A, M, expect)
}
