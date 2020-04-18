package main

import "testing"

func runSample(t *testing.T, m, n int, grid [][]int, expect int64) {
	res := solve(m, n, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", m, n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 3, 3
	grid := [][]int{
		{0, 9, 9},
		{0, 6, 6},
		{0, 9, 6},
	}
	var expect int64 = 39
	runSample(t, m, n, grid, expect)
}

func TestSample2(t *testing.T) {
	m, n := 2, 2
	grid := [][]int{
		{1, 1},
		{1, 1},
	}
	var expect int64 = 4
	runSample(t, m, n, grid, expect)
}
func TestSample3(t *testing.T) {
	m, n := 1, 4
	grid := [][]int{
		{1, 2, 3, 4},
	}
	var expect int64 = 9
	runSample(t, m, n, grid, expect)
}

func TestSample4(t *testing.T) {
	m, n := 4, 3
	grid := [][]int{
		{4, 1, 4},
		{1, 6, 1},
		{1, 2, 1},
		{3, 2, 3},
	}
	var expect int64 = 18
	runSample(t, m, n, grid, expect)
}

func TestSample5(t *testing.T) {
	m, n := 4, 4
	grid := [][]int{
		{83, 43, 69, 43},
		{35, 45, 39, 33},
		{9, 76, 39, 55},
		{22, 88, 58, 60},
	}
	var expect int64 = 503
	runSample(t, m, n, grid, expect)
}

func TestSample6(t *testing.T) {
	m, n := 4, 4
	grid := [][]int{
		{34, 70, 66, 56},
		{94, 89, 22, 67},
		{22, 9, 26, 79},
		{30, 78, 73, 63},
	}
	var expect int64 = 554
	runSample(t, m, n, grid, expect)
}
