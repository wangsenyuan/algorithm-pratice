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
		{1, -1},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, -1, -1, -1},
		{-1, 1, 1, -1},
		{1, 1, 1, -1},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, -1, 1, 1},
		{-1, 1, -1, 1},
		{1, -1, 1, 1},
	}
	expect := false
	runSample(t, grid, expect)
}
