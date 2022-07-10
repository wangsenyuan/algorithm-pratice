package p2333

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, k1, k2 int, expect int64) {
	res := minSumSquareDiff(nums1, nums2, k1, k2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 10, 20, 19}
	k1, k2 := 0, 0
	var expect int64 = 579
	runSample(t, nums1, nums2, k1, k2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 4, 10, 12}
	nums2 := []int{5, 8, 6, 9}
	k1, k2 := 1, 1
	var expect int64 = 43
	runSample(t, nums1, nums2, k1, k2, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{1, 4, 10, 12}
	nums2 := []int{5, 8, 6, 9}
	k1, k2 := 10, 5
	var expect int64 = 0
	runSample(t, nums1, nums2, k1, k2, expect)
}

func TestSample4(t *testing.T) {
	nums1 := []int{7, 11, 4, 19, 11, 5, 6, 1, 8}
	nums2 := []int{4, 7, 6, 16, 12, 9, 10, 2, 10}
	k1, k2 := 3, 6
	var expect int64 = 27
	runSample(t, nums1, nums2, k1, k2, expect)
}

func TestSample5(t *testing.T) {
	nums1 := []int{19, 18, 19, 18, 18, 19, 19}
	nums2 := []int{1, 0, 1, 0, 0, 1, 1}
	k1, k2 := 10, 33
	var expect int64 = 985
	runSample(t, nums1, nums2, k1, k2, expect)
}
