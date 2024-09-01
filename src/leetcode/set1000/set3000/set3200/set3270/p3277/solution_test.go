package p3277

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := maximumSubarrayXor(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 8, 4, 32, 16, 1}
	queries := [][]int{
		{0, 2},
		{1, 4},
		{0, 5},
	}
	expect := []int{12, 60, 60}
	runSample(t, nums, queries, expect)
}
