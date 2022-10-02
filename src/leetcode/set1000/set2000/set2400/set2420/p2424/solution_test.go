package p2424

import "testing"

func runSample(t *testing.T, nums1, nums2 []int, diff int, expect int64) {
	res := numberOfPairs(nums1, nums2, diff)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{3, 2, 5}
	nums2 := []int{2, 2, 1}
	diff := 1
	var expect int64 = 3
	runSample(t, nums1, nums2, diff, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{3, -1}
	nums2 := []int{-2, 2}
	diff := -1
	var expect int64
	runSample(t, nums1, nums2, diff, expect)
}
