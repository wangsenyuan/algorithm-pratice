package p2449

import "testing"

func runSample(t *testing.T, nums []int, target []int, expect int64) {
	res := makeSimilar(nums, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 12, 6}
	target := []int{2, 14, 10}
	var expect int64 = 2
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 5}
	target := []int{4, 1, 3}
	var expect int64 = 1
	runSample(t, nums, target, expect)
}
