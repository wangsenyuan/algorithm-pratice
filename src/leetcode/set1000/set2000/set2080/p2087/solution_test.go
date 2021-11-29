package p2087

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := countPyramids(grid)

	if res != expect {
		t.Errorf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 1, 1, 0}, {1, 1, 1, 1},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 1}, {1, 1, 1},
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 1},
	}
	expect := 13
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{0, 1, 1},
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	expect := 1
	runSample(t, grid, expect)
}
