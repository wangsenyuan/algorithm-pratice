package main

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	x, y := solve(nums)
	if expect != (x > 0) {
		t.Fatalf("Sample expect %t, but got %d %d", expect, x, y)
	}

	if !expect {
		return
	}
	a, b, c, d := nums[0], nums[1], nums[2], nums[3]

	if x <= a || x > c || y <= b || y > d {
		t.Fatalf("Sample result %d %d out of bounds", x, y)
	}
	if int64(x*y)%int64(a*b) != 0 {
		t.Fatalf("Sample result %d*%d not divides %d * %d", x, y, a, b)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1024, 729, 373248, 730}
	expect := true
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1024, 729, 373247, 730}
	expect := false
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{5040, 40320, 40319, 1000000000}
	expect := true
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{999999999, 999999999, 1000000000, 1000000000}
	expect := false
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{268435456, 268435456, 1000000000, 1000000000}
	expect := true
	runSample(t, nums, expect)
}
