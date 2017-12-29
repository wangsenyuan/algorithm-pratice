package main

import "testing"

func runSample(t *testing.T, n int, grid []string, expect int) {
	mat := make([][]byte, n)
	for i := 0; i < n; i++ {
		mat[i] = []byte(grid[i])
	}
	res := solve(n, mat)
	if res != expect {
		t.Errorf("sample: %d %v, should give %d, but got %d", n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	grid := []string{
		"...",
		".##",
		"*#.",
	}
	runSample(t, n, grid, -1)
}

func TestSample2(t *testing.T) {
	n := 3
	grid := []string{
		"..*",
		"...",
		"...",
	}
	runSample(t, n, grid, 4)
}

func TestSample3(t *testing.T) {
	n := 3
	grid := []string{
		"..*",
		"*..",
		"...",
	}
	runSample(t, n, grid, 6)
}

func TestSample4(t *testing.T) {
	n := 4
	grid := []string{
		"....",
		".#.*",
		".#*.",
		"**#.",
	}
	runSample(t, n, grid, 16)
}
