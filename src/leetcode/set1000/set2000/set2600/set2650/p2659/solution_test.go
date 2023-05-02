package p2659

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := countOperationsToEmptyArray(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, -1}
	var expect int64 = 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4, 3}
	var expect int64 = 5
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	var expect int64 = 3
	runSample(t, nums, expect)
}
