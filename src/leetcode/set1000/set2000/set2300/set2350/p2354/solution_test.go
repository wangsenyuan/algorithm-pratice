package p2354

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := countExcellentPairs(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	var expect int64 = 5
	runSample(t, nums, k, expect)
}
