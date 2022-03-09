package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 7}
	expect := []int{1, 1, 2, 2, 4}
	for i := 0; i < len(nums); i++ {
		runSample(t, nums[i], expect[i])
	}
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 4)
}
