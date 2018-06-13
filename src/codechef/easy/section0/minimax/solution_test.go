package main

import "testing"

func runSample(t *testing.T, n int, grid [][]int, expect int) {
	res := solve(n, grid)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	grid := [][]int{
		{10, 20, 30},
		{20, 10, 30},
		{10, 5, 35},
	}
	expect := 1
	runSample(t, n, grid, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	grid := [][]int{
		{10, 20, 10},
		{20, 10, 5},
		{30, 30, 35},
	}
	expect := 0
	runSample(t, n, grid, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	grid := [][]int{
		{1, 1, 3, 4},
		{5, 1, 1, 8},
		{9, 10, 1, 1},
		{1, 14, 15, 1},
	}
	expect := 2
	runSample(t, n, grid, expect)
}
