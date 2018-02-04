package p778

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := swimInWater(grid)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 2},
		{1, 3},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6},
	}
	expect := 16
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{3, 2},
		{0, 1},
	}
	expect := 3
	runSample(t, grid, expect)
}
