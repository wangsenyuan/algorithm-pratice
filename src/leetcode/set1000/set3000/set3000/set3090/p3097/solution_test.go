package p3097

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minimumSubarrayLength(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 12, 8, 55, 88, 2}
	k := 85
	expect := 1
	runSample(t, nums, k, expect)
}
