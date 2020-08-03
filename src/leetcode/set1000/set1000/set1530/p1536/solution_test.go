package p1536

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minSwaps(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 0, 1}, {1, 1, 0}, {1, 0, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0,1,1,0}, {0,1,1,0}, {0,1,1,0}, {0,1,1,0},
	}
	expect := -1
	runSample(t, grid, expect)
}