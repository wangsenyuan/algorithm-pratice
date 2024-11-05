package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 2}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 0, 2}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 0, 1, 4, 2}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{0, 0, 0, 0, 1, 2, 4}
	expect := 7
	runSample(t, nums, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{2, 0}
	expect := 2
	runSample(t, nums, expect)
}
