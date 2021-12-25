package main

import "testing"

func runSample(t *testing.T, s string, Q [][]int, expect []string) {
	solver := NewSolver(s)

	for i, cur := range Q {
		L, R, K := cur[0], cur[1], cur[2]
		res := solver.solve(L, R, K)

		if res != expect[i] {
			t.Fatalf("Sample result %s, not correct %v", res, expect)
		}
	}
}

func TestSample1(t *testing.T) {
	S := "EFDAABCBD"
	Q := [][]int{
		{1, 9, 6},
		{6, 8, 1},
		{4, 7, 3},
		{1, 2, 1},
	}
	expect := []string{
		"FEDAABBCD",
		"",
		"CAAB",
		"FE",
	}

	runSample(t, S, Q, expect)
}
