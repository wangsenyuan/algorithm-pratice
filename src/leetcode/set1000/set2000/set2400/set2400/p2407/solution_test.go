package p2407

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := lengthOfLIS(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 2, 1, 4, 3, 4, 5, 8, 15}
	k := 3
	expect := 5
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 4, 5, 1, 8, 12, 4, 7}
	k := 5
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 5}
	k := 1
	expect := 1
	runSample(t, nums, k, expect)
}
