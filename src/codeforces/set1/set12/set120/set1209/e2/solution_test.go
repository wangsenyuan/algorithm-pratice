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
		{2, 5, 7},
		{4, 2, 4},
	}
	expect := 12
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{4, 1, 5, 2, 10, 4},
		{8, 6, 6, 4, 9, 10},
		{5, 4, 9, 5, 8, 7},
	}
	expect := 29
	runSample(t, grid, expect)
}
