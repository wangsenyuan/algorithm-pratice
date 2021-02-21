package p1769

import "testing"

func runSample(t *testing.T, nums []int, mul []int, expect int) {
	res := maximumScore(nums, mul)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	mul := []int{3, 2, 1}
	expect := 14
	runSample(t, nums, mul, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-5, -3, -3, -2, 7, 1}
	mul := []int{-10, -5, 3, 4, 6}
	expect := 102
	runSample(t, nums, mul, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	mul := []int{1, 2, 3}
	expect := 14
	runSample(t, nums, mul, expect)
}
