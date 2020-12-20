package p1695

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maximumUniqueSubarray(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 2, 4, 5, 6}
	expect := 17
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	expect := 8
	runSample(t, nums, expect)
}
