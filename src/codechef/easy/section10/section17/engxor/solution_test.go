package main

import "testing"

func runSample(t *testing.T, n int, A []int, queries []int, expect []int) {
	solver := NewSolver(n, A)

	for i := 0; i < len(queries); i++ {
		a, _ := solver.Ask(queries[i])

		if a != expect[i] {
			t.Errorf("Sample expect %v, but got %d at step %d", expect, a, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{4, 2, 15, 9, 8, 8}
	queries := []int{3}
	expect := []int{2}

	runSample(t, n, A, queries, expect)
}
