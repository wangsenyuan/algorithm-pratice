package main

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := solve(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"11..2",
		"#..22",
		"#.323",
		".#333",
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"1#2#3",
	}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		".2..",
		"...3",
		".1#.",
	}
	expect := 2
	runSample(t, grid, expect)
}
