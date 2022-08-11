package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 10, 20}
	expect := 3
	runSample(t, nums, expect)
}
