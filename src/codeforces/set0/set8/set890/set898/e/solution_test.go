package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{12, 14, 30, 4}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 0}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{120, 110, 23, 34, 25, 45}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{121, 56, 78, 81, 45, 100, 1, 0, 54, 78}
	expect := 0
	runSample(t, nums, expect)
}
