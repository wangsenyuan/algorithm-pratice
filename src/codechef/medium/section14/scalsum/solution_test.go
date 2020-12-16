package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, W []int, Q [][]int, expect []uint) {
	solver := NewSolver(n, E, W)

	for i := 0; i < len(Q); i++ {
		u, v := Q[i][0], Q[i][1]
		res := solver.Query(u, v)
		if res != expect[i] {
			t.Fatalf("query(%d %d) expect %d, but got %d", u, v, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	W := []int{5, 4, 3, 2, 1}
	Q := [][]int{
		{2, 3},
		{4, 5},
	}
	expect := []uint{37, 43}

	runSample(t, n, E, W, Q, expect)
}
