package p1798

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxScore(nums)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	runSample(t, nums, 14)
}
