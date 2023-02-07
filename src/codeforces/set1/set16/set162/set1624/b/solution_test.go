package main

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := solve(nums)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 1}
	expect := false
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 5, 30}
	expect := true
	runSample(t, nums, expect)
}
