package p1568

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minDays(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 0, 1, 1},
	}
	expect := 1
	runSample(t, grid, expect)
}
