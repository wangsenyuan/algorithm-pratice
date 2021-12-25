package main

import "testing"

func runSample(t *testing.T, m, n, x int, grid []string, expect int) {
	res := solve(m, n, x, grid)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", m, n, x, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n, x := 3, 3, 0
	grid := []string{
		".*.",
		"*.*",
		".*.",
	}
	expect := -1
	runSample(t, m, n, x, grid, expect)
}

func TestSample2(t *testing.T) {
	m, n, x := 3, 3, 1
	grid := []string{
		".*.",
		"*.*",
		".*.",
	}
	expect := 3
	runSample(t, m, n, x, grid, expect)
}

func TestSample3(t *testing.T) {
	m, n, x := 3, 3, 2
	grid := []string{
		".*.",
		"*.*",
		".*.",
	}
	expect := 4
	runSample(t, m, n, x, grid, expect)
}
