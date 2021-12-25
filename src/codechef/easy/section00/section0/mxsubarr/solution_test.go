package main

import "testing"

func runSample(t *testing.T, n int, nums []int, expect int64) {
	res := solve(n, nums)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	nums := []int{3, 4}
	runSample(t, n, nums, 1)
}

func TestSample2(t *testing.T) {
	n := 3
	nums := []int{1, 2, 2}
	runSample(t, n, nums, 3)
}

func TestSample3(t *testing.T) {
	n := 5
	nums := []int{1, 1, 2, 2, 2}
	runSample(t, n, nums, 18)
}

func TestSample4(t *testing.T) {
	n := 1
	nums := []int{1}
	runSample(t, n, nums, 1)
}
