package p1909

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := canBeIncreasing(nums)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 10, 5, 7}
	expect := true
	runSample(t, nums, expect)
}
