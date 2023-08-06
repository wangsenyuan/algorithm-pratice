package p2806

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, x int, expect int) {
	res := minimumTime(nums1, nums2, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 3}
	x := 4
	expect := 3
	runSample(t, nums1, nums2, x, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{3, 3, 3}
	x := 4
	expect := -1
	runSample(t, nums1, nums2, x, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{8, 2, 3}
	nums2 := []int{1, 4, 2}
	x := 13
	expect := 0
	runSample(t, nums1, nums2, x, expect)
}
