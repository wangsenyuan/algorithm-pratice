package main

import "testing"

func runSample(t *testing.T, n int, nums []int, expect int) {
	res := solve(n, nums)
	if res != expect {
		t.Errorf("sample %d %v, should give %d, but got %d", n, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	nums := []int{2, 1, 3}
	expect := 19
	runSample(t, n, nums, expect)
}
