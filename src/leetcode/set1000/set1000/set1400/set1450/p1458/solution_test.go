package p1458

import "testing"

func runSample(t *testing.T, nums1, nums2 []int, expect int) {
	res := maxDotProduct(nums1, nums2)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", nums1, nums2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{2, 1, -2, 5}
	nums2 := []int{3, 0, -6}
	expect := 18
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{3, -2}
	nums2 := []int{2, -6, 7}
	expect := 21
	runSample(t, nums1, nums2, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{-1, -1}
	nums2 := []int{1, 1}
	expect := -1
	runSample(t, nums1, nums2, expect)
}

func TestSample4(t *testing.T) {
	nums1 := []int{6, -1, 8, -3, -7}
	nums2 := []int{0, -2, 4, -6, -4, 10, 1, 5}
	expect := 110
	runSample(t, nums1, nums2, expect)
}
