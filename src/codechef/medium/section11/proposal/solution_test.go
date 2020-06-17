package main

import "testing"

func runSample(t *testing.T, n int, nums []int, expect int) {
	res := solve(n, nums)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	nums := []int{2 ,3, 4, 6, 8}
	expect := 1

	runSample(t, n, nums, expect)
}