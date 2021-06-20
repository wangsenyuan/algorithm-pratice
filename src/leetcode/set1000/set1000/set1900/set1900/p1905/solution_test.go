package p1905

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := minDifference(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 4, 8}
	queries := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {0, 3},
	}
	expect := []int{2, 1, 4, 1}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 5, 2, 2, 7, 10}
	queries := [][]int{
		{2, 3}, {0, 2}, {0, 5}, {3, 5},
	}
	expect := []int{-1, 1, 1, 3}
	runSample(t, nums, queries, expect)
}
