package main

import "testing"

func runSample(t *testing.T, n, w int, A []int, Q [][]int, expect []int64) {
	solver := solve(n, w, A)

	for i := 0; i < len(Q); i++ {
		p, c := Q[i][0], Q[i][1]
		res := solver.Update(p, c)

		if res != expect[i] {
			t.Fatalf("Sample expect %v, but got (%d %d)", expect, i, res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, w := 4, 2
	A := []int{1, 1, 1, 1}

	Q := [][]int{
		{1, 1},
		{1, 2},
		{2, 2},
	}
	expect := []int64{3, 2, 2}
	runSample(t, n, w, A, Q, expect)
}
