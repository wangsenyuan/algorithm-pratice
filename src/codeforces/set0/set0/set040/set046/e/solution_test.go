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
		{-1, 2},
		{1, 3},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{2, -6, 2},
		{-2, -2, 1},
	}
	expect := -4
	runSample(t, grid, expect)
}
