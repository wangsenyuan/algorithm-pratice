package p1705

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect []int) {
	res := findBall(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1},
	}
	expect := []int{1, -1, -1, -1, -1}
	runSample(t, grid, expect)
}
