package lcp14

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := splitArray(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 3, 2, 3, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 5, 7}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{326614, 489921}
	expect := 1
	runSample(t, nums, expect)
}
