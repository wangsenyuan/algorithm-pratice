package p1969

import "testing"

func runSample(t *testing.T, row, col int, cells [][]int, expect int) {
	res := latestDayToCross(row, col, cells)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	row := 2
	col := 2
	cells := [][]int{
		{1, 1}, {2, 1}, {1, 2}, {2, 2},
	}
	expect := 2
	runSample(t, row, col, cells, expect)
}

func TestSample2(t *testing.T) {
	row := 2
	col := 2
	cells := [][]int{
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}
	expect := 1
	runSample(t, row, col, cells, expect)
}
