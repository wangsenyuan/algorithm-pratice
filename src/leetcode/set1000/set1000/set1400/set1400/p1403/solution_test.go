package p1403

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := minSubsequence(nums)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 3, 10, 9, 8}
	expect := []int{10, 9}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 4, 7, 6, 7}
	expect := []int{7, 7, 6}
	runSample(t, nums, expect)
}
