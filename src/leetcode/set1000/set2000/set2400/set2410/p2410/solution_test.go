package p2410

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := smallestSubarrays(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 0, 2, 1, 3}
	expect := []int{3, 3, 2, 2, 1}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2}
	expect := []int{2, 1}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 0, 5, 6, 3, 2}
	expect := []int{4, 3, 2, 2, 1, 1}
	runSample(t, nums, expect)
}
