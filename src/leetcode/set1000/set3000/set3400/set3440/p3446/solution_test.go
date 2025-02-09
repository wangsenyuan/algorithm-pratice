package p3446

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect [][]int) {
	res := sortMatrix(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}
	expect := [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}
	runSample(t, grid, expect)
}
