package main

import "testing"

func runSample(t *testing.T, white [][]int, black [][]int, queries [][]int, expect []int) {
	solver := NewSolver(len(white), white, len(black), black)

	for i := 0; i < len(queries); i++ {
		res := solver.Query(queries[i])

		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	white := [][]int{
		{0, 0},
		{2, 0},
		{2, 3},
		{0, 3},
	}
	black := [][]int{
		{1, 1},
		{1, 2},
	}
	queries := [][]int{
		{4, 3, 2, 1},
		{1, 4, 2},
		{1, 3, 2},
	}
	expect := []int{2, 1, 1}

	runSample(t, white, black, queries, expect)
}
