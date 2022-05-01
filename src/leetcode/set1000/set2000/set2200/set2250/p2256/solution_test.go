package p2256

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minimumAverageDifference(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 5, 3, 9, 5, 3}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0}
	expect := 0
	runSample(t, nums, expect)
}
