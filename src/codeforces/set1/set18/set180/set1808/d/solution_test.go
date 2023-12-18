package main

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := solve(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 8, 2, 5, 2, 8, 6}
	k := 5
	expect := 4
	runSample(t, nums, k, expect)
}
