package main

import "testing"

func runSample(t *testing.T, nums []int, s string, expect int) {
	res := solve(nums, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	s := "+ + *"
	expect := 3
	runSample(t, nums, s, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 2, 2}
	s := "* * +"
	expect := 8
	runSample(t, nums, s, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	s := "* + +"
	expect := 9
	runSample(t, nums, s, expect)
}
