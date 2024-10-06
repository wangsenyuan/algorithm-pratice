package p3312

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries []int64, expect []int) {
	res := gcdValues(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 4}
	queries := []int64{0, 2, 2}
	expect := []int{1, 2, 2}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 4, 2, 1}
	queries := []int64{5, 3, 1, 0}
	expect := []int{4, 2, 1, 1}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2}
	queries := []int64{0, 0}
	expect := []int{2, 2}
	runSample(t, nums, queries, expect)
}
