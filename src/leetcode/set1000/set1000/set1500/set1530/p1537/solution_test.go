package p1537

import "testing"

func runSample(t *testing.T, nums1, nums2 []int, expect int)  {
	res := maxSum(nums1, nums2)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", nums1, nums2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{2,4,5,8,10}
	nums2 := []int{4,6,8,9}
	expect := 30
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1,3,5,7,9}
	nums2 := []int{3,5,100}
	expect := 109
	runSample(t, nums1, nums2, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{1,2,3,4,5}
	nums2 := []int{6,7,8,9,10}
	expect := 40
	runSample(t, nums1, nums2, expect)
}