package p2755

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := continuousSubarrays(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 4, 2, 4}
	var expect int64 = 8
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3}
	var expect int64 = 6
	runSample(t, nums, expect)
}
