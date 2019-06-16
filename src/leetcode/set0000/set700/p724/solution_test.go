package p724

import "testing"

func TestSample1(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	if res != 3 {
		t.Errorf("pivot index of %v should be 3, but got %d", nums, res)
	}
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3}
	res := pivotIndex(nums)
	if res != -1 {
		t.Errorf("pivot index of %v should be -1, but got %d", nums, res)
	}
}
