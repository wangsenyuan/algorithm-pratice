package p2501

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, queries []int, expect []int) {
	res := maxPoints(grid, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 2, 3}, {2, 5, 7}, {3, 5, 1},
	}
	queries := []int{5, 6, 2}
	expect := []int{5, 8, 1}
	runSample(t, grid, queries, expect)
}
