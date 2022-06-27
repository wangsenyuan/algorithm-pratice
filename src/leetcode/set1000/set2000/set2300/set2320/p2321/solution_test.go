package p2321

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := maximumsSplicedArray(nums1, nums2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{60, 60, 60}
	nums2 := []int{10, 90, 10}
	expect := 210
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{20, 40, 20, 70, 30}
	nums2 := []int{50, 20, 50, 40, 20}
	expect := 220
	runSample(t, nums1, nums2, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{7, 11, 13}
	nums2 := []int{1, 1, 1}
	expect := 31
	runSample(t, nums1, nums2, expect)
}
