package p769

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxChunksToSorted(nums)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 3, 2, 1, 0}
	runSample(t, nums, 1)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 2, 3, 4}
	runSample(t, nums, 4)
}
