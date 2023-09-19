package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 5, 2, 3}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{17, 8, 19, 45}
	expect := 63
	runSample(t, nums, expect)
}
