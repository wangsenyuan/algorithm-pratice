package reversepairs

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := reversePairs(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{7, 5, 6, 4}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 2, 3, 1}
	expect := 4
	runSample(t, nums, expect)
}
