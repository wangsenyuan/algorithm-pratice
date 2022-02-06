package p2163

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := int(minimumDifference(nums))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{7, 9, 5, 8, 1, 3}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23}
	expect := -14
	runSample(t, nums, expect)
}
