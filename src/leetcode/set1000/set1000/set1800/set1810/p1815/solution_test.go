package p1815

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countNicePairs(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{42, 11, 1, 97}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{13, 10, 35, 24, 76}
	expect := 4
	runSample(t, nums, expect)
}
