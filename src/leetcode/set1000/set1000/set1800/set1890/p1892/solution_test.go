package p1892

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := largestMagicSquare(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{6, 9, 10, 5, 6, 5, 6},
		{9, 8, 1, 6, 2, 6, 8},
		{9, 3, 5, 7, 6, 5, 3},
		{6, 4, 9, 2, 7, 8, 5},
	}
	expect := 3
	runSample(t, grid, expect)
}
