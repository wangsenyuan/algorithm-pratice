package main

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-1, 1, -1, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-1, 2, -3, 2, -1}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, -1}
	expect := false
	runSample(t, A, expect)
}
