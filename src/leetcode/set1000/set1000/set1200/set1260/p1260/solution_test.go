package p1260

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, k int, expect [][]int) {
	res := shiftGrid(grid, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %d, expect %v, but got %v", grid, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1}, {2}, {3}, {4}, {7}, {6}, {5},
	}
	k := 23
	runSample(t, grid, k, grid)
}
