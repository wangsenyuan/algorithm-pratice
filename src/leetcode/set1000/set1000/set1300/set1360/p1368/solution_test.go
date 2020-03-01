package p1368

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minCost(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 3}, {3, 2, 2}, {1, 1, 4},
	}
	expect := 0
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 2}, {4, 3},
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{2, 2, 2}, {2, 2, 2},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := [][]int{
		{4},
	}
	expect := 0
	runSample(t, grid, expect)
}
