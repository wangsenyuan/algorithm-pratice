package p980

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := uniquePathsIII(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2},
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 1}, {2, 0},
	}
	expect := 0
	runSample(t, grid, expect)
}
