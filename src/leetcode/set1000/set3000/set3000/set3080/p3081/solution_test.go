package p3081

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int64) {
	res := unmarkedSumArray(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 2, 1, 2, 3, 1}
	queries := [][]int{
		{1, 2}, {3, 3}, {4, 2},
	}
	expect := []int64{8, 3, 0}
	runSample(t, nums, queries, expect)
}
