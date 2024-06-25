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
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	expect := 7
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0},
	}
	expect := 8
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
	}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample6(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
	}
	expect := 12
	runSample(t, grid, expect)
}
