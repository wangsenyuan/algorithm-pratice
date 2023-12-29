package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 1, 3, 4, 6}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1000000000, 123456789, 987654321, 998244353, 500000004}
	expect := 10590032
	runSample(t, nums, expect)
}
