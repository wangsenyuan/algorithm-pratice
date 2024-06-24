package main

import (
	"testing"
)

func runSample(t *testing.T, grid [][]int, k int, d int, expect int) {
	res := solve(grid, k, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0},
		{0, 1, 2, 3, 2, 1, 2, 3, 3, 2, 0},
		{0, 1, 2, 3, 5, 5, 5, 5, 5, 2, 0},
	}
	k, d := 1, 4
	expect := 4
	runSample(t, grid, k, d, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 3, 3, 0},
		{0, 2, 1, 0},
		{0, 1, 2, 0},
		{0, 3, 3, 0},
	}
	k, d := 2, 1
	expect := 8
	runSample(t, grid, k, d, expect)
}
