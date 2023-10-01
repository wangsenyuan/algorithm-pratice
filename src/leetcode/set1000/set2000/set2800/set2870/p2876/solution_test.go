package p2876

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maximumTripletValue(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{12, 6, 1, 2, 7}
	var expect int64 = 77
	runSample(t, nums, expect)
}
