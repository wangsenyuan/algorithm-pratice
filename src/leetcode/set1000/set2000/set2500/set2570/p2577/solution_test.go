package p2577

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minimumTime(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 1, 3, 2},
		{5, 1, 2, 5},
		{4, 3, 8, 6},
	}
	expect := 7
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 2, 4},
		{3, 2, 1},
		{1, 0, 4},
	}
	expect := -1
	runSample(t, grid, expect)
}
