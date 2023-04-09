package p2617

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minimumVisitedCells(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0},
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{3, 4, 2, 1}, {4, 2, 1, 1}, {2, 1, 1, 0}, {3, 4, 1, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}
