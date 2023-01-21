package main

import "testing"

func runSample(t *testing.T, n int, m int, panel [][]int, symbols [][]int, x int, y int, expect int) {
	res := solve(n, m, panel, symbols, x, y)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 5
	panel := [][]int{
		{9, 0, 1, 1, 0},
		{0, 0, 3, 2, 0},
		{1, 1, 1, 3, 0},
		{1, 1, 1, 3, 0},
		{0, 1, 2, 0, 3},
	}
	symbols := [][]int{
		{-1, 1, -1, 3, -1},
		{-1, -1, -1, 0, -1},
		{-1, -1, -1, -1, -1},
		{-1, 2, 3, -1, -1},
		{-1, -1, -1, -1, 2},
	}
	x, y := 4, 2
	expect := 35
	runSample(t, n, m, panel, symbols, x, y, expect)
}
