package p2877

import "testing"

func runSample(t *testing.T, nums []int, target int, expect int) {
	res := minSizeSubarray(nums, target)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1}
	target := 83
	expect := 53
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2}
	target := 72
	expect := 48
	runSample(t, nums, target, expect)
}
