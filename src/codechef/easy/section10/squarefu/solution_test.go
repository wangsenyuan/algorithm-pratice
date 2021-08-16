package main

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := solve(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 12}
	var expect int64 = 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	var expect int64 = 5
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 3, 7}
	var expect int64 = 3
	runSample(t, nums, expect)
}
