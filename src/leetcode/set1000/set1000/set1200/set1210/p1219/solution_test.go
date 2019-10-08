package p1219

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := getMaximumGold(grid)

	if res != expect {
		t.Errorf("Sample %v, epxect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 6, 0}, {5, 8, 7}, {0, 9, 0},
	}

	expect := 24
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20},
	}

	expect := 28
	runSample(t, grid, expect)
}
