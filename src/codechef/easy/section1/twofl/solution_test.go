package main

import "testing"

func runSample(t *testing.T, n int, m int, grid [][]int, expect int) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	grid := [][]int{
		{1, 1, 2, 3, 1},
		{3, 1, 2, 5, 2},
		{5, 2, 1, 5, 6},
		{1, 3, 1, 2, 1},
	}
	expect := 10
	runSample(t, n, m, grid, expect)
}
