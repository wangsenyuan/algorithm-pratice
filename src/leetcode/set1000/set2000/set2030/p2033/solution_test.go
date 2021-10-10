package p2033

import "testing"

func runSample(t *testing.T, grid [][]int, x int, expect int) {
	res := minOperations(grid, x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{2, 4},
		{6, 8},
	}
	x := 2
	expect := 4
	runSample(t, grid, x, expect)
}
