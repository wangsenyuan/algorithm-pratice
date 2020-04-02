package main

import "testing"

func runSample(t *testing.T, N int, M int, E [][]int, expect int) {
	res := solve(N, M, E)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 4
	E := [][]int{
		{1, 2, 10, 15},
		{1, 3, 10, 16},
		{2, 4, 1, 1},
		{3, 5, 10, 15},
	}
	expect := 50
	runSample(t, N, M, E, expect)
}
