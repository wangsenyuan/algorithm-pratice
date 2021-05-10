package p1855

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := maxDistance(nums1, nums2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{55, 30, 5, 4, 2}
	b := []int{100, 20, 10, 10, 5}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 2}
	b := []int{10, 10, 1}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{30, 29, 19, 5}
	b := []int{25, 25, 25, 25, 25}
	expect := 2
	runSample(t, a, b, expect)
}
