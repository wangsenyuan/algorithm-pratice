package p1775

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := minOperations(nums1, nums2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 1, 2, 2, 2, 2}
	expect := 3
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{6, 6}
	nums2 := []int{1}
	expect := 3
	runSample(t, nums1, nums2, expect)
}
