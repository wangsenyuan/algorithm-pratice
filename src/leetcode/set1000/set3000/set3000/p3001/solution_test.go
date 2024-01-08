package p3001

import "testing"

func runSample(t *testing.T, nums1, nums2 []int, expect int) {
	res := maximumSetSize(nums1, nums2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{2, 3, 2, 3, 2, 3}
	expect := 5
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 1, 2, 2, 3, 3}
	nums2 := []int{4, 4, 5, 5, 6, 6}
	expect := 6
	runSample(t, nums1, nums2, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{3, 5}
	nums2 := []int{5, 3}
	expect := 2
	runSample(t, nums1, nums2, expect)
}

func TestSample4(t *testing.T) {
	nums1 := []int{1, 2, 1, 2}
	nums2 := []int{1, 1, 1, 1}
	expect := 2
	runSample(t, nums1, nums2, expect)
}
