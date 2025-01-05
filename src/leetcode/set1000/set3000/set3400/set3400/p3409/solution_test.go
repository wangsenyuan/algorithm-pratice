package p3409

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestSubsequence(nums)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{16, 6, 3}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 5, 3, 4, 2, 1}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 20, 10, 19, 10, 20}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 8, 8, 8, 1}
	expect := 4
	runSample(t, nums, expect)
}
