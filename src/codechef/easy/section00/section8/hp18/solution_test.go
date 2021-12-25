package main

import "testing"

func runSample(t *testing.T, N, a, b int, nums []int, expect bool) {
	res := solve(N, a, b, nums)
	if expect != res {
		t.Errorf("Sample %d %d %d %v, expect %t, but got %t", N, a, b, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, a, b := 5, 3, 2
	nums := []int{1, 2, 3, 4, 5}
	runSample(t, N, a, b, nums, false)
}

func TestSample2(t *testing.T) {
	N, a, b := 5, 2, 4
	nums := []int{1, 2, 3, 4, 5}
	runSample(t, N, a, b, nums, true)
}
