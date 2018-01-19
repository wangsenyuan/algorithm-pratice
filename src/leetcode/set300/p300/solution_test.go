package p300

import "testing"

func TestSample(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expect := 4
	res := lengthOfLIS(nums)
	if res != expect {
		t.Errorf("sample %v, should give %d, but got %d", nums, expect, res)
	}
}
