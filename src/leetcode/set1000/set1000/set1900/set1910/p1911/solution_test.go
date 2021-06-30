package p1911

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maxAlternatingSum(nums)

	if res != expect {
		t.Errorf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 2, 5, 3}
	var expect int64 = 7
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6, 7, 8}
	var expect int64 = 8
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{6, 2, 1, 2, 4, 5}
	var expect int64 = 10
	runSample(t, nums, expect)
}
