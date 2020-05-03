package p1438

import "testing"

func runSample(t *testing.T, nums []int, limit int, expect int) {
	res := longestSubarray(nums, limit)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", nums, limit, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 2, 4, 7}
	limit := 4
	expect := 2
	runSample(t, nums, limit, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 1, 2, 4, 7, 2}
	limit := 5
	expect := 4
	runSample(t, nums, limit, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit := 0
	expect := 3
	runSample(t, nums, limit, expect)
}
