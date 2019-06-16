package p750

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := countCornerRectangles(grid)
	if res != expect {
		t.Errorf("sample: %v, should give %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, [][]int{{1, 0, 0, 1, 0}, {0, 0, 1, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 0, 1}}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, [][]int{{1, 1, 1, 1}}, 0)
}
