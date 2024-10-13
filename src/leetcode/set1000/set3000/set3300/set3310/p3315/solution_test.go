package p3315

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := minBitwiseArray(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 5, 7}
	expect := []int{-1, 1, 4, 3}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{11, 13, 31}
	expect := []int{9, 12, 15}
	runSample(t, nums, expect)
}
