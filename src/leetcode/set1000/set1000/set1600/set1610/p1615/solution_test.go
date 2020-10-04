package p1615

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := specialArray(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 6, 7, 7, 0}
	expect := -1
	runSample(t, nums, expect)
}
