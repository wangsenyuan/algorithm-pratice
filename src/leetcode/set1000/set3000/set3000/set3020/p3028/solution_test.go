package p3028

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid [][]int, threshold int, expect [][]int) {
	res := resultGrid(grid, threshold)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v,but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	image := [][]int{
		{5, 6, 7, 10},
		{8, 9, 10, 10},
		{11, 12, 13, 10},
	}
	threshold := 3
	expect := [][]int{
		{9, 9, 9, 9},
		{9, 9, 9, 9},
		{9, 9, 9, 9},
	}
	runSample(t, image, threshold, expect)
}

func TestSample2(t *testing.T) {
	image := [][]int{
		{10, 20, 30}, {15, 25, 35}, {20, 30, 40}, {25, 35, 45},
	}
	threshold := 12
	expect := [][]int{
		{25, 25, 25}, {27, 27, 27}, {27, 27, 27}, {30, 30, 30},
	}
	runSample(t, image, threshold, expect)
}

func TestSample3(t *testing.T) {
	image := [][]int{
		{0, 7, 0},
		{4, 6, 3},
		{2, 4, 5},
	}
	threshold := 5
	expect := [][]int{
		{0, 7, 0},
		{4, 6, 3},
		{2, 4, 5},
	}
	runSample(t, image, threshold, expect)
}
