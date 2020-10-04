package p1595

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxProductPath(grid)

	if res != expect {
		t.Errorf("Sample %v expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{-1, -2, -3},
		{-2, -3, -3},
		{-3, -3, -2},
	}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, -2, 1},
		{1, -2, 1},
		{3, -4, 1},
	}
	expect := 8
	runSample(t, grid, expect)
}
