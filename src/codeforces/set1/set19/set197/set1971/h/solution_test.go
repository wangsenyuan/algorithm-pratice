package main

import "testing"

func runSample(t *testing.T, grid [][]int, expect bool) {
	res := solve(grid)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, -2, -3, -2},
		{-4, 4, -1, -3},
		{1, 2, -2, 4},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2, 3, 4, 5},
		{-2, 3, -4, -5, -1},
		{3, -5, 1, 2, 2},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{-1, -2},
		{2, -2},
	}
	expect := false
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{1, 3, -6, 2, 5, 2},
		{1, 3, -2, -3, -6, -5},
		{-2, -1, -3, 2, 3, 1},
	}
	expect := false
	runSample(t, grid, expect)
}
