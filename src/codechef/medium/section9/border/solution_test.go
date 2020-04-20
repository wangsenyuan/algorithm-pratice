package main

import "testing"

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	solver := NewSolver(s)

	for i, query := range queries {
		res := solver.Query(query[0], query[1])

		if res != expect[i] {
			t.Fatalf("Sample fail at %d, with expect %d, but got %d", i, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "abcab"
	queries := [][]int{
		{2, 1},
		{5, 1},
		{5, 2},
		{5, 3},
	}
	expect := []int{2, 2, 5, -1}
	runSample(t, s, queries, expect)
}
