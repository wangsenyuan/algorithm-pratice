package main

import "testing"

func runSample(t *testing.T, nums []int, c int, expect int) {
	res := solve(c, nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{9, 9, 9, 9, 9}
	c := 9
	expect := 5
	runSample(t, nums, c, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 2, 6}
	c := 2
	expect := 2
	runSample(t, nums, c, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 4, 3, 4, 3, 3, 4, 1, 1, 3, 4, 1, 3, 3, 4, 1, 3, 1, 2, 3, 1, 4, 2, 2, 3, 4, 2, 4, 2}
	c := 3
	expect := 12
	runSample(t, nums, c, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 3, 2, 2, 3, 2, 1, 2, 2, 1, 3, 1, 1, 3, 3, 3, 3, 2}
	c := 2
	expect := 11
	runSample(t, nums, c, expect)
}
