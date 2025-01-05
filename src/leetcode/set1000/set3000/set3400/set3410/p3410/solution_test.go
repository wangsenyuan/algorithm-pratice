package p3410

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxSubarraySum(nums)
	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{-3, 2, -2, -1, 3, -2, 3}
	expect := 7
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 10
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-31, -23, -47}
	expect := -23
	runSample(t, nums, expect)
}
