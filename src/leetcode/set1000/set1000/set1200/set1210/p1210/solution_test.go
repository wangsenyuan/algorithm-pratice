package p1210

import "testing"

func runMinimumMoves(t *testing.T, grid [][]int, expect int) {
	res := minimumMoves(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}
func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
	}
	runMinimumMoves(t, grid, 11)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 0, 0, 1},
		{1, 1, 1, 0, 0, 0},
	}
	runMinimumMoves(t, grid, 9)
}
