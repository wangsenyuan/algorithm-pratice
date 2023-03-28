package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{60, 12, 30, 24, 58, 33, 18, 25, 16}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 5, 13, 9, 7, 25, 21, 6}
	expect := 3
	runSample(t, nums, expect)
}
