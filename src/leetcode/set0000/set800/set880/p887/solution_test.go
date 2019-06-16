package p887

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := projectionArea(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{2},
	}

	expect := 5

	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2}, {3, 4},
	}

	expect := 17

	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 0}, {0, 2},
	}

	expect := 8

	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}

	expect := 14

	runSample(t, grid, expect)
}
