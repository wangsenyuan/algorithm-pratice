package p1914

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, k int, expect [][]int) {
	res := rotateGrid(grid, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{40, 10}, {30, 20}}
	expect := [][]int{
		{10, 20}, {40, 30},
	}
	k := 1
	runSample(t, grid, k, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	expect := [][]int{
		{3, 4, 8, 12}, {2, 11, 10, 16}, {1, 7, 6, 15}, {5, 9, 13, 14},
	}
	k := 2
	runSample(t, grid, k, expect)
}
