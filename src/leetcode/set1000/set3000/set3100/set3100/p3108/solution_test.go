package p3108

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestMonotonicSubarray(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 9, 7, 1}
	expect := 3
	runSample(t, nums, expect)
}
