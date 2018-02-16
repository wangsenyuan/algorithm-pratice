package main

import "testing"

func runSample(t *testing.T, n int, nums []int, expect int) {
	res := solve(n, nums)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	nums := []int{1, 2, 3, 4, 5, 6}
	runSample(t, n, nums, 1)
}
