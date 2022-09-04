package p2401

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestNiceSubarray(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 8, 48, 10}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 5, 11, 13}
	expect := 1
	runSample(t, nums, expect)
}
