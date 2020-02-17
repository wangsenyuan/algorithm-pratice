package p1351

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := countNegatives(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3},
	}

	expect := 8
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, -1}, {-1, -1},
	}

	expect := 3
	runSample(t, grid, expect)
}
