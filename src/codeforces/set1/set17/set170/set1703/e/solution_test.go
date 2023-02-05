package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := solve(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"010",
		"110",
		"010",
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"11100",
		"11011",
		"01011",
		"10011",
		"11000",
	}
	expect := 9
	runSample(t, grid, expect)
}
