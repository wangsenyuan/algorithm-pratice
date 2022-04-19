package p2245

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxTrailingZeros(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{4, 3, 2}, {7, 6, 1}, {8, 8, 8},
	}
	expect := 0
	runSample(t, grid, expect)
}
