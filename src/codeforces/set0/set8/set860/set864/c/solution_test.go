package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums[0], nums[1], nums[2], nums[3])

	if res != expect {
		t.Fatalf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 1, 1}
	expect := -1
	runSample(t, nums, expect)
}
