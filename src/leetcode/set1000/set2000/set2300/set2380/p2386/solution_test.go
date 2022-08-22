package p2386

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := kSum(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, -2}
	k := 5
	var expect int64 = 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, -2, 3, 4, -10, 12}
	k := 16
	var expect int64 = 10
	runSample(t, nums, k, expect)
}
