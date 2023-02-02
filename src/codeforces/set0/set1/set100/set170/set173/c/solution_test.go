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
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	expect := 17
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{1, 1, 1},
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{-3, 2, 0, 1, 5, -1},
		{4, -1, 2, -3, 0, 1},
		{-5, 1, 2, 4, 1, -2},
		{0, -2, 1, 3, -1, 2},
		{3, 1, 4, -3, -2, 0},
		{-1, 2, -1, 3, 1, 2},
	}
	expect := 13
	runSample(t, grid, expect)
}
