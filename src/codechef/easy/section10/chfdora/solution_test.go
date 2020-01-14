package main

import "testing"

func runSample(t *testing.T, N, M int, grid [][]int, expect int64) {
	res := solve(N, M, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 3, 3
	grid := [][]int{
		{2, 1, 2},
		{1, 1, 1},
		{2, 1, 2},
	}
	expect := int64(10)
	runSample(t, N, M, grid, expect)
}
