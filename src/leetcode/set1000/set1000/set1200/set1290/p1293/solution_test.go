package p1293

import "testing"

func runSample(t *testing.T, grid [][]int, k, expect int) {
	res := shortestPath(grid, k)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 1, 1}, {1, 1, 1}, {1, 0, 0},
	}
	k := 1
	expect := -1
	runSample(t, grid, k, expect)
}
