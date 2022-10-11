package p2435

import (
	"testing"
)

func runSample(t *testing.T, grid [][]int, k int, expect int) {
	res := numberOfPaths(grid, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{5, 2, 4},
		{3, 0, 5},
		{0, 7, 2},
	}
	k := 3
	expect := 2
	runSample(t, grid, k, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 0},
	}
	k := 5
	expect := 1
	runSample(t, grid, k, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{7, 3, 4, 9},
		{2, 3, 6, 2},
		{2, 3, 7, 0},
	}
	k := 1
	expect := 10
	runSample(t, grid, k, expect)
}
