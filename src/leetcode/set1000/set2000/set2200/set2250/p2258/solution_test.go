package p2258

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maximumMinutes(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{0, 2, 0, 0, 0, 0, 0}, {0, 0, 0, 2, 2, 1, 0}, {0, 2, 0, 0, 1, 2, 0}, {0, 0, 2, 2, 2, 0, 2}, {0, 0, 0, 0, 0, 0, 0},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0}, {0, 1, 2, 0}, {0, 2, 0, 0},
	}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 0, 0}, {2, 2, 0}, {1, 2, 0},
	}
	expect := 1000000000
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{0,2,0,0,1}, {0,2,0,2,2}, {0,2,0,0,0}, {0,0,2,2,0}, {0,0,0,0,0},
	}
	expect := 0
	runSample(t, grid, expect)
}
