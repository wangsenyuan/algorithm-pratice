package main

import "testing"

func runSample(t *testing.T, n int, nums []int, expect int) {
	res := solve(n, nums)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	nums := []int{0, 0}
	expect := 1
	runSample(t, n, nums, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	nums := []int{3, 2, 1}
	expect := 2
	runSample(t, n, nums, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	nums := []int{4, 8, 6, 1, 5, 2}
	expect := 4
	runSample(t, n, nums, expect)
}
