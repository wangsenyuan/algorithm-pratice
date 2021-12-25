package main

import "testing"

func runSample(t *testing.T, n, k int, nums []int, expect int) {
	res := solve(n, k, nums)
	if res != expect {
		t.Errorf("sample: %d %d %v, should give %d, but got %d", n, k, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	nums := []int{0, 1, 2}
	expect := 2
	runSample(t, n, k, nums, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 3
	nums := []int{0, 1, 2, 3, 4, 0, 1, 2, 5, 3}
	expect := 379
	runSample(t, n, k, nums, expect)
}
