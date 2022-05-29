package p2288

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := totalSteps(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 14, 4, 14, 13, 2, 6, 13}
	expect := 3
	runSample(t, nums, expect)
}
