package main

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := solve(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{4, 2, 11, 10},
		{6, 7, 1, 9},
		{5, 12, 3, 8},
	}
	expect := 9
	runSample(t, grid, expect)
}
