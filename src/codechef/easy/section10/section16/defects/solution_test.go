package main

import "testing"

func runSample(t *testing.T, n, m int, grid [][]int, expect int) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 6
	grid := [][]int{
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
	}

	expect := 2
	runSample(t, n, m, grid, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 3
	grid := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	expect := 0
	runSample(t, n, m, grid, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 4
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
	}

	expect := 1
	runSample(t, n, m, grid, expect)
}

func TestSample4(t *testing.T) {
	n, m := 1, 5
	grid := [][]int{
		{1, 0, 1, 0, 1},
	}

	expect := 2
	runSample(t, n, m, grid, expect)
}
