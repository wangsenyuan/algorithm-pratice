package main

import "testing"

func runSample(t *testing.T, n int, grid []string, expect int) {
	res := solve(n, grid)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	grid := []string{
		"100000",
		"100000",
		"100000",
		"100000",
		"010010",
		"001100",
	}
	expect := 2
	runSample(t, n, grid, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	grid := []string{
		"0011",
		"1100",
		"1110",
		"0001",
	}
	expect := 0
	runSample(t, n, grid, expect)
}
