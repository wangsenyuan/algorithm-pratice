package main

import "testing"

func runSample(t *testing.T, grid []string, src []int, dst []int, expect bool) {
	res := solve(len(grid), len(grid[0]), grid, src, dst)

	if expect != res {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"X...XX",
		"...XX.",
		".X..X.",
		"......",
	}
	src := []int{1, 6}
	dst := []int{2, 2}
	expect := true
	runSample(t, grid, src, dst, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		".X..",
		"...X",
		"X.X.",
		"....",
		".XX.",
	}
	src := []int{5, 3}
	dst := []int{1, 1}
	expect := false
	runSample(t, grid, src, dst, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"..X.XX.",
		".XX..X.",
		"X...X..",
		"X......",
	}
	src := []int{2, 2}
	dst := []int{1, 6}
	expect := true
	runSample(t, grid, src, dst, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		".XX",
		"...",
		".X.",
		".X.",
		"...",
	}
	src := []int{1, 3}
	dst := []int{4, 1}
	expect := true
	runSample(t, grid, src, dst, expect)
}
