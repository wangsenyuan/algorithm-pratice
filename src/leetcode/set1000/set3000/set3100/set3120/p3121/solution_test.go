package p3121

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minimumOperations(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 0, 2}, {1, 0, 2},
	}
	expect := 0
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 1}, {0, 0, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1}, {2}, {3},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{0, 2, 6, 0, 4, 8, 0, 5},
		{3, 3, 9, 9, 3, 0, 3, 8},
		{8, 3, 0, 5, 1, 6, 3, 6},
		{0, 9, 2, 0, 6, 6, 9, 7},
		{5, 8, 9, 3, 7, 1, 3, 4},
		{0, 2, 0, 7, 9, 2, 6, 2},
		{2, 4, 7, 1, 4, 7, 8, 2},
		{3, 8, 0, 4, 0, 7, 6, 4},
	}
	expect := 46
	runSample(t, grid, expect)
}
