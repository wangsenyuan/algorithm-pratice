package lcp49

import "testing"

func runSample(t *testing.T, nums []int64, expect int64) {
	res := ringGame(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int64{94, 95, 94, 95, 1, 95, 1, 5, 1}
	var expect int64 = 90
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int64{5, 4, 6, 2, 7}
	var expect int64 = 4
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int64{31, 2, 31, 29, 29, 1, 31}
	var expect int64 = 29
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int64{6, 2, 4, 71, 71, 3, 71, 1, 71, 71}
	var expect int64 = 65
	runSample(t, nums, expect)
}
