package p2302

import "testing"

func runSample(t *testing.T, nums []int, k int64, expect int64) {
	res := countSubarrays(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 4, 3, 5}
	var k int64 = 10
	var expect int64 = 6
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1}
	var k int64 = 5
	var expect int64 = 5
	runSample(t, nums, k, expect)
}
