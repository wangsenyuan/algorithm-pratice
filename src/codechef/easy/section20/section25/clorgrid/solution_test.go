package main

import "testing"

func runSample(t *testing.T, n int, m int, grid []string, expect bool) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	grid := []string{
		"..#",
		"...",
		"...",
	}
	expect := true
	runSample(t, n, m, grid, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	grid := []string{
		"...",
		".#.",
		"...",
	}
	expect := false
	runSample(t, n, m, grid, expect)
}

