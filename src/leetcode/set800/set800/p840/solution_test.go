package p840

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := numMagicSquaresInside(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}

	runSample(t, grid, 1)
}
