package p749

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := containVirus(grid)
	if res != expect {
		t.Errorf("sample: %v, should give %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	expect := 10
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0},
	}
	expect := 13
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expect := 56
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 0, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0, 0, 1},
		{1, 0, 1, 1, 0, 1, 0, 1, 0, 1},
	}
	expect := 38
	runSample(t, grid, expect)
}
