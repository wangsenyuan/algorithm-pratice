package p2007

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := findOriginalArray(nums)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 4, 2, 6, 8}
	expect := []int{1, 3, 4}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 3, 0, 1}
	runSample(t, nums, nil)
}

func TestSample3(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	expect := []int{0, 0}
	runSample(t, nums, expect)
}
