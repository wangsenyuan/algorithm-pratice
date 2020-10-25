package p5547

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, l []int, r []int, expect []bool) {
	res := checkArithmeticSubarrays(nums, l, r)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	expect := []bool{true, false, true}
	runSample(t, nums, l, r, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
	l := []int{0, 1, 6, 4, 8, 7}
	r := []int{4, 4, 9, 7, 9, 10}
	expect := []bool{false, true, false, false, true, true}
	runSample(t, nums, l, r, expect)
}
