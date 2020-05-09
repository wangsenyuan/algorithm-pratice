package main

import "testing"

func runSample(t *testing.T, n, m int, grid []string, expect int) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	grid := []string{
		".#.",
		"###",
		"##.",
	}
	expect := 1
	runSample(t, n, m, grid, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 2
	grid := []string{
		"##",
		".#",
		".#",
		"##",
	}
	expect := -1
	runSample(t, n, m, grid, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 1
	grid := []string{
		".",
	}
	expect := 0
	runSample(t, n, m, grid, expect)
}

func TestSample4(t *testing.T) {
	n, m := 2, 1
	grid := []string{
		".",
		"#",
	}
	expect := -1
	runSample(t, n, m, grid, expect)
}
