package p2652

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, x int, expect []int) {
	res := getSubarrayBeauty(nums, k, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -1, -3, -2, 3}
	k := 3
	x := 2
	expect := []int{-1, -2, -2}
	runSample(t, nums, k, x, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	k := 2
	x := 2
	expect := []int{-1, -2, -3, -4}
	runSample(t, nums, k, x, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-3, 1, 2, -3, 0, -3}
	k := 2
	x := 1
	expect := []int{-3, 0, -3, -3, -3}
	runSample(t, nums, k, x, expect)
}
