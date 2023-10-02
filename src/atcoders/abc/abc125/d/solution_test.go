package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-10, 5, -4}
	expect := 19
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, -4, -8, -11, 3}
	expect := 30
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1000000000, 1000000000, -1000000000, 1000000000, -1000000000, 0, 1000000000, -1000000000, 1000000000, -1000000000, 1000000000}
	expect := 10000000000
	runSample(t, a, expect)
}
