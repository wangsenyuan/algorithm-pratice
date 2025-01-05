package p3414

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, intervals [][]int, expect []int) {
	res := maximumWeight(intervals)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	intervals := [][]int{
		{1, 3, 2}, {4, 5, 2}, {1, 5, 5}, {6, 9, 3}, {6, 7, 1}, {8, 9, 1},
	}
	expect := []int{2, 3}
	runSample(t, intervals, expect)
}

func TestSample2(t *testing.T) {
	intervals := [][]int{
		{5, 8, 1}, {6, 7, 7}, {4, 7, 3}, {9, 10, 6}, {7, 8, 2}, {11, 14, 3}, {3, 5, 5},
	}
	expect := []int{1, 3, 5, 6}
	runSample(t, intervals, expect)
}

func TestSample3(t *testing.T) {
	intervals := [][]int{
		{7, 9, 37}, {20, 24, 36}, {10, 11, 28}, {23, 25, 19}, {4, 21, 31}, {1, 3, 9}, {11, 18, 27}, {25, 25, 40},
	}
	expect := []int{0, 1, 2, 7}
	runSample(t, intervals, expect)
}
