package main

import (
	"testing"
)

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 3, 4, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 2}
	expect := -1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 3, 2, 1, 2, 3, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	expect := 0
	runSample(t, nums, expect)
}
