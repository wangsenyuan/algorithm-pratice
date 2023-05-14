package p2680

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := maximumOr(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{12, 9}
	k := 1
	var expect int64 = 30
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{8, 1, 2}
	k := 2
	var expect int64 = 35
	runSample(t, nums, k, expect)
}
