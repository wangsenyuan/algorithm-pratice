package snailsort

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, expect []int) {
	res := Snail(grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expect := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{}
	//var expect []int
	runSample(t, grid, []int{})
}

func TestSample3(t *testing.T) {
	grid := [][]int{{1}}
	expect := []int{1}
	runSample(t, grid, expect)
}
