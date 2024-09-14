package main

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := solve(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{3, 2, 1},
		{1, 2, 3},
		{4, 5, 6},
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	expect := 0
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 6, 3, 4},
		{5, 10, 7, 8},
		{9, 2, 11, 12},
	}
	expect := 2
	runSample(t, grid, expect)
}
