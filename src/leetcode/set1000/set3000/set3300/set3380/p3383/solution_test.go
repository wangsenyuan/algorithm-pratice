package p3383

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxSubarraySum(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	k := 1
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, -2, -3, -4, -5}
	k := 4
	expect := -10
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-5, 1, 2, -3, 4}
	k := 2
	expect := 4
	runSample(t, a, k, expect)
}
