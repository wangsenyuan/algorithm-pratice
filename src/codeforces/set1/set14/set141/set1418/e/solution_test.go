package main

import "testing"

func runSample(t *testing.T, n int, D []int, queries [][]int, expect []int) {
	solver := NewSolver(n, D)

	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		a, b := cur[0], cur[1]
		res := solver.solve(a, b)
		if res != expect[i] {
			t.Fatalf("Sample result %d, not correct, expect %d", res, expect[i])
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	D := []int{1, 3, 1}
	queries := [][]int{
		{2, 1},
		{1, 2},
	}
	expect := []int{665496237, 1}
	runSample(t, n, D, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	D := []int{4, 2, 6}
	queries := [][]int{
		{3, 1},
		{1, 2},
		{2, 3},
	}
	expect := []int{0, 8, 665496236}
	runSample(t, n, D, queries, expect)
}
