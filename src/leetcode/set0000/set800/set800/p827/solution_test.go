package p827

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := largestIsland(grid)

	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 0}, {0, 1}}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{1, 1}, {0, 1}}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{{1, 1}, {1, 1}}
	expect := 4
	runSample(t, grid, expect)
}
