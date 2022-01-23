package p2146

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, pricing []int, start []int, k int, expect [][]int) {
	res := highestRankedKItems(grid, pricing, start, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1},
	}
	pricings := []int{2, 5}
	start := []int{0, 0}
	k := 3
	expect := [][]int{{0, 1}, {1, 1}, {2, 1}}
	runSample(t, grid, pricings, start, k, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1},
	}
	pricings := []int{2, 3}
	start := []int{2, 3}
	k := 2
	expect := [][]int{{2, 1}, {1, 2}}
	runSample(t, grid, pricings, start, k, expect)
}
