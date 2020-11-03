package main

import "testing"

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []int) {
	solver := NewSolver(n, A)

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		res := solver.Update(cur[0], cur[1])

		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 1, 2, 5, 2}
	queries := [][]int{
		{1, 3},
		{4, 2},
	}
	expect := []int{5, 3}
	runSample(t, n, A, queries, expect)
}
