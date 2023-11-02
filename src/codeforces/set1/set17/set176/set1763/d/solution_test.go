package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 3, 2, 3}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 3, 3, 2}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{20, 6, 15, 8, 17}
	expect := 4788
	runSample(t, nums, expect)
}
