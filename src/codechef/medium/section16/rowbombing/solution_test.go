package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := solve(grid)

	if res != expect {
		t.Errorf("Smaple expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"..",
	}
	expect := 0
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		".##.",
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		".",
		"#",
		"#",
		".",
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		".#",
		"##",
		"##",
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := []string{
		".##",
		"#.#",
		"##.",
	}
	expect := 1
	runSample(t, grid, expect)
}
