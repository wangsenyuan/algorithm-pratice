package chapter4

import "testing"

func runQuickSelect(t *testing.T, nums []int, k int, expect int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	res := QuickSelect(nums, k)
	if res != expect {
		t.Errorf("QuickSelect Sample %v, %d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestQuickSelect1(t *testing.T) {
	nums := []int{4, 1, 10, 8, 7, 12, 9, 2, 15}
	k := 5
	expect := 8
	runQuickSelect(t, nums, k, expect)
}
