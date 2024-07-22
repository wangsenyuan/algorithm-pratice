package p3224

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maximumScore(grid)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0},
		{0, 1, 0, 0, 0},
		{5, 0, 0, 3, 0},
		{0, 0, 0, 0, 2},
	}
	expect := 11
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{10, 9, 0, 0, 15},
		{7, 1, 0, 8, 0},
		{5, 20, 0, 11, 0},
		{0, 0, 0, 1, 2},
		{8, 12, 1, 10, 3},
	}
	expect := 94
	runSample(t, grid, expect)
}
