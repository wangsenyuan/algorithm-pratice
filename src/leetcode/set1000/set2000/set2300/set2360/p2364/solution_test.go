package p2364

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := countBadPairs(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 1, 3, 3}
	var expect int64 = 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	var expect int64
	runSample(t, nums, expect)
}
