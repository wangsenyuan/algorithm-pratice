package p2541

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, k int, expect int64) {
	res := minOperations(nums1, nums2, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{4, 3, 1, 4}
	nums2 := []int{1, 3, 7, 1}
	k := 3

	var expect int64 = 2
	runSample(t, nums1, nums2, k, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{3, 8, 5, 2}
	nums2 := []int{2, 4, 1, 6}
	k := 1

	var expect int64 = -1
	runSample(t, nums1, nums2, k, expect)
}
