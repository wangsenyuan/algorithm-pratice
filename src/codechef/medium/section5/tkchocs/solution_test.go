package main

import "testing"

func runSample(t *testing.T, N int, grid [][]int, expect int) {
	res := solve(N, grid)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 2
	grid := [][]int{
		{1},
		{1, 1},
	}
	expect := 3
	runSample(t, N, grid, expect)
}

func TestSample2(t *testing.T) {
	N := 3
	grid := [][]int{
		{1},
		{2, 5},
		{1, 3, 1},
	}
	expect := 11
	runSample(t, N, grid, expect)
}
