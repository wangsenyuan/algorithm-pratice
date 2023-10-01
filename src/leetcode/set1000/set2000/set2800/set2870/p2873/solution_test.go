package p2873

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxSubarrays(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 0, 2, 0, 1, 2}
	expect := 3
	runSample(t, nums, expect)
}
