package p1878

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect []int) {
	res := getBiggestThree(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5},
	}
	expect := []int{228, 216, 211}
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	expect := []int{20, 9, 8}
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{20, 17, 9, 13, 5, 2, 9, 1, 5}, {14, 9, 9, 9, 16, 18, 3, 4, 12}, {18, 15, 10, 20, 19, 20, 15, 12, 11}, {19, 16, 19, 18, 8, 13, 15, 14, 11}, {4, 19, 5, 2, 19, 17, 7, 2, 2},
	}
	expect := []int{107, 103, 102}
	runSample(t, grid, expect)
}
