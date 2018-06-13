package main

import "testing"

func runSample(t *testing.T, n, m int, grid [][]int, expect int) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("sample %d %d %v, expect %d, but got %d", n, m, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	grid := [][]int{
		{1, 2, 3},
		{4, -10, 5},
		{6, 7, 8},
	}

	expect := -10

	runSample(t, n, m, grid, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	grid := [][]int{
		{1, -10, 3},
		{-13, -12, -11},
		{7, 8, 9},
	}

	expect := -46

	runSample(t, n, m, grid, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	grid := [][]int{
		{1, -2, 3},
		{-4, 5, -6},
		{7, -8, 9},
	}

	expect := -15

	runSample(t, n, m, grid, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 4
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, -7, 8},
		{-9, -10, -11, -12},
		{13, 14, -15, 16},
	}

	expect := -64

	runSample(t, n, m, grid, expect)
}
