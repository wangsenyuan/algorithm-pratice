package main

import "testing"

func runSample(t *testing.T, grid []string, x int, expect int) {
	res := solve(grid, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"1.#2",
		"#..#",
		"*.#.",
	}
	x := 3
	expect := 8
	runSample(t, grid, x, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"2.#2",
		"#..#",
		"*.#.",
	}
	x := 5
	expect := -1
	runSample(t, grid, x, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"2.#2",
		"#.3#",
		"*.#.",
	}
	x := 5
	expect := 6
	runSample(t, grid, x, expect)
}
