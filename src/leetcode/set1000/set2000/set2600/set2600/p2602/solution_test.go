package p2602

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries []int, expect []int64) {
	res := minOperations(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 6, 8}
	queries := []int{1, 5}
	expect := []int64{14, 10}
	runSample(t, nums, queries, expect)
}
