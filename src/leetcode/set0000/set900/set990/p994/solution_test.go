package p994

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := orangesRotting(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
	}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 2},
	}
	expect := 0
	runSample(t, grid, expect)
}
