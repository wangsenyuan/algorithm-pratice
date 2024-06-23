package p3197

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minimumSum(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 1},
	}

	expect := 3
	runSample(t, grid, expect)
}
