package main

import "testing"

func runSample(t *testing.T, S string, queries [][]int, expect []int) {
	solver := NewSolver(S)

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		solver.Flip(cur[0], cur[1])
		res := solver.GetResult()

		if res != expect[i] {
			t.Fatalf("Fail at step %d, expect %d, but got %d", i, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	S := "000"
	queries := [][]int{
		{3, 3},
		{2, 3},
		{3, 3},
		{2, 3},
		{1, 3},
	}
	expect := []int{
		1,
		1,
		2,
		1,
		4,
	}
	runSample(t, S, queries, expect)
}
