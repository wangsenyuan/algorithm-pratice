package p2304

import "testing"

func runSample(t *testing.T, grid [][]int, moveCost [][]int, expect int) {
	res := minPathCost(grid, moveCost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{5, 1, 2}, {4, 0, 3},
	}
	moveCost := [][]int{
		{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2},
	}
	expect := 6
	runSample(t, grid, moveCost, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{5, 3}, {4, 0}, {2, 1},
	}
	moveCost := [][]int{
		{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3},
	}
	expect := 17
	runSample(t, grid, moveCost, expect)
}
