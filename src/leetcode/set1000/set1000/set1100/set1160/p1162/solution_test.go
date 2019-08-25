package p1162

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxDistance(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample(t *testing.T) {
	grid := [][]int{
		{1, 0, 1}, {0, 0, 0}, {1, 0, 1},
	}
	expect := 2
	runSample(t, grid, expect)
}
