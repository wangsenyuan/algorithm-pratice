package p1463

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := cherryPickup(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 8, 7, 10, 9, 10, 0, 9, 6}, {8, 7, 10, 8, 7, 4, 9, 6, 10}, {8, 1, 1, 5, 1, 5, 5, 1, 2}, {9, 4, 10, 8, 8, 1, 9, 5, 0}, {4, 3, 6, 10, 9, 2, 4, 8, 10}, {7, 3, 2, 8, 3, 3, 5, 9, 8}, {1, 2, 6, 5, 6, 2, 0, 10, 0},
	}
	expect := 96
	runSample(t, grid, expect)
}
