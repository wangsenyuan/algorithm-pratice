package p1793

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumScore(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 3, 7, 4, 5}
	k := 3
	expect := 15
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 5, 4, 5, 4, 1, 1, 1}
	k := 0
	expect := 20
	runSample(t, nums, k, expect)
}
