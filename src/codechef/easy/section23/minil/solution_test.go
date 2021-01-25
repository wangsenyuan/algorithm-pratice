package main

import "testing"

func runSample(t *testing.T, m, n int, grid []string, expect int) {
	res := solve(m, n, grid)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 2, 2
	grid := []string{
		"..",
		"*.",
	}
	expect := 1
	runSample(t, m, n, grid, expect)
}

func TestSample2(t *testing.T) {
	m, n := 2, 4
	grid := []string{
		"*..*",
		"*.*.",
	}
	expect := 2
	runSample(t, m, n, grid, expect)
}
