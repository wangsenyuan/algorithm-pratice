package p3021

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maximumLength(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 4, 1, 2, 2}
	expect := 3
	runSample(t, nums, expect)
}
