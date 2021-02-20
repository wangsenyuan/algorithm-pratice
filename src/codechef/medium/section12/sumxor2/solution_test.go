package main

import "testing"

func runSample(t *testing.T, n int, A []int, Q []int, expect []int) {
	solver := NewSolver(n, A)

	for i, m := range Q {
		cur := solver.Query(m)

		if cur != expect[i] {
			t.Fatalf("Sample expect %v, but got %d", expect, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 3, 5, 2}
	Q := []int{1, 2}
	expect := []int{11, 34}
	runSample(t, n, A, Q, expect)
}
