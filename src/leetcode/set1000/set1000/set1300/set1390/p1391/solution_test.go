package p1391

import "testing"

func runSample(t *testing.T, grid [][]int, expect bool) {
	res := hasValidPath(grid)

	if res != expect {
		t.Errorf("sample %v, expect %t, but got %t", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{4, 1}, {6, 1},
	}
	expect := true

	runSample(t, grid, expect)
}
