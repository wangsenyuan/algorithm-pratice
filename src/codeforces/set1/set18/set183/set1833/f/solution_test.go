package main

import "testing"

func runSample(t *testing.T, nums []int, m int, expect int) {
	res := solve(nums, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 10, 10, 9, 6, 11, 7}
	m := 4
	expect := 5
	runSample(t, nums, m, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 2, 2, 3, 6}
	m := 3
	expect := 2
	runSample(t, nums, m, expect)
}
