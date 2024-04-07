package p3107

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minOperationsToMakeMedianK(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample %v %d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 5, 6, 8, 5}
	k := 4
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 5, 6, 8, 5}
	k := 7
	expect := 3
	runSample(t, nums, k, expect)
}
