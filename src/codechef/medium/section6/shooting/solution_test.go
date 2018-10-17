package main

import "testing"

func runSample(t *testing.T, n, m int, grid []string, expect bool) {
	res := solve(n, m, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, m, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	grid := []string{
		"E.",
		"L.",
	}
	expect := true
	runSample(t, n, m, grid, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 4
	grid := []string{
		"E.EL",
		"LL..",
	}
	expect := true
	runSample(t, n, m, grid, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	grid := []string{
		"EE.",
		"L.L",
		"L..",
	}
	expect := false
	runSample(t, n, m, grid, expect)
}
