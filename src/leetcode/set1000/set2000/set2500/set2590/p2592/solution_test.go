package p2592

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maximizeGreatness(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{42, 8, 75, 28, 35, 21, 13, 21}
	expect := 6
	runSample(t, nums, expect)
}
