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
	nums := []int{1, 1, 2, 2}
	expect := true
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 4, 5, 7}
	expect := true
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{8, 9, 15, 18}
	expect := true
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{12, 21, 14, 24}
	expect := false
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{36, 60, 48, 66}
	expect := false
	runSample(t, nums, expect)
}
