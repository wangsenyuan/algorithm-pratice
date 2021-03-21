package p1803

import "testing"

func runSample(t *testing.T, nums []int, low, high int, expect int) {
	res := countPairs(nums, low, high)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 2, 7}
	low, high := 2, 6
	expect := 6
	runSample(t, nums, low, high, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{9, 8, 4, 2, 1}
	low, high := 5, 14
	expect := 8
	runSample(t, nums, low, high, expect)
}
