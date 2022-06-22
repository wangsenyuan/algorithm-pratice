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
		{7, 3, 2},
		{4, 1, 5},
		{6, 8, 9},
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{9, 8, 5},
		{2, 4, 1},
		{3, 7, 6},
	}
	expect := -1
	runSample(t, grid, expect)
}
