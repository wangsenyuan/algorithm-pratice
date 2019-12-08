package p1283

import "testing"

func runSample(t *testing.T, nums []int, threshold, expect int) {
	res := smallestDivisor(nums, threshold)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", nums, threshold, expect, res)
	}
}

func TestSample(t *testing.T) {
	runSample(t, []int{1, 2, 5, 9}, 6, 5)
}
