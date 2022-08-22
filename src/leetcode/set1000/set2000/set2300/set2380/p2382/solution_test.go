package p2382

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, removes []int, expect []int64) {
	res := maximumSegmentSum(nums, removes)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 5, 6, 1}
	rms := []int{0, 3, 2, 4, 1}
	expect := []int64{14, 7, 2, 2, 0}
	runSample(t, nums, rms, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 11, 1}
	rms := []int{3, 2, 1, 0}
	expect := []int64{16, 5, 3, 0}
	runSample(t, nums, rms, expect)
}
