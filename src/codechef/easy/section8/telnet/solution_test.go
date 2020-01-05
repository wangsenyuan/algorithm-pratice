package main

import "testing"

func runSample(t *testing.T, n int, k int, grid [][]int, expect int) {
	res := solve(n, k, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	grid := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}
	expect := 2
	runSample(t, n, k, grid, expect)
}
