package p2584

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := findValidSplit(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 7, 8, 15, 3, 5}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 7, 15, 8, 3, 5}
	expect := -1
	runSample(t, nums, expect)
}
