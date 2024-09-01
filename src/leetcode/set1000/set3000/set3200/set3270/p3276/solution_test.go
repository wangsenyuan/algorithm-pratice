package p3276

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxScore(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{8, 7, 6},
		{8, 3, 2},
	}
	expect := 15
	runSample(t, grid, expect)
}
