package p3074

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minimumOperationsToWriteY(grid)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 2, 2}, {1, 1, 0}, {0, 1, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}
