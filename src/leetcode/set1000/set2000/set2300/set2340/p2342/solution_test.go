package p2342

import (
	"testing"
)

func runSample(t *testing.T, nums []int, expect int) {
	res := maximumSum(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 12, 19, 14}
	expect := -1
	runSample(t, nums, expect)
}
