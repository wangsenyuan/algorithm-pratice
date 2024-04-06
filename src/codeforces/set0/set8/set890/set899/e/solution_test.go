package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 5, 5, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 3, 4, 1, 5}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 4, 4, 2, 2, 100, 100, 100}
	expect := 3
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{10, 10, 50, 10, 50, 50}
	expect := 4
	runSample(t, a, expect)
}
