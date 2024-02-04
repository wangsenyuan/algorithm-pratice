package p3025

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumSubarraySum(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 1
	expect := 11
	runSample(t, nums, k, expect)
}
