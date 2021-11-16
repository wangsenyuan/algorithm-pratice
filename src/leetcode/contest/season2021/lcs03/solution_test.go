package lcs03

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := largestArea(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"110", "231", "221",
	}
	expect := 1
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"11111100000", "21243101111", "21224101221", "11111101111",
	}
	expect := 3
	runSample(t, grid, expect)
}
