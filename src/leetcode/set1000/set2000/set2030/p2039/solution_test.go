package p2039

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, k int64, expect int64) {
	res := kthSmallestProduct(nums1, nums2, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{2, 5}
	nums2 := []int{3, 4}
	k := int64(2)
	expect := int64(8)
	runSample(t, nums1, nums2, k, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{-4, -2, 0, 3}
	nums2 := []int{2, 4}
	k := int64(6)
	expect := int64(0)
	runSample(t, nums1, nums2, k, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{-2, -1, 0, 1, 2}
	nums2 := []int{-3, -1, 2, 4, 5}
	k := int64(3)
	expect := int64(-6)
	runSample(t, nums1, nums2, k, expect)
}
