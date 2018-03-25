package p807

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxIncreaseKeepingSkyline(grid)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	runSample(t, grid, 35)
}
