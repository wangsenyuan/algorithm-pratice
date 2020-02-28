package main

import "testing"

func runSample(t *testing.T, n int, grid [][]int, expect int) {
	res := solve(n, grid)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	grid := [][]int{
		{1, 3, 2},
		{3, 4, 3},
		{5, 3, 7},
	}
	expect := 4
	runSample(t, n, grid, expect)
}
