package p2560

import "testing"

func runSample(t *testing.T, grid [][]int, expect bool) {
	res := isPossibleToCutPath(grid)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0},
		{0, 0, 1, 0, 1, 1},
	}
	expect := true
	runSample(t, grid, expect)
}
