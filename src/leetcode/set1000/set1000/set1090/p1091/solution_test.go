package p1091

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := shortestPathBinaryMatrix(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}
