package main

import "testing"

func runSample(t *testing.T, n int, m int, mat [][]int, queries [][]int, expect []bool) {
	solver := NewSolver(n, m, mat)

	for i, q := range queries {
		res := solver.Change(q[0], q[1], q[2])
		if res != expect[i] {
			t.Fatalf("result %t, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 3
	mat := [][]int{
		{1, 4, 3},
		{7, 1, 4},
	}
	qs := [][]int{
		{1, 2, 4},
		{1, 1, 5},
		{2, 2, 5},
	}
	expect := []bool{true, false, true}

	runSample(t, n, m, mat, qs, expect)
}
