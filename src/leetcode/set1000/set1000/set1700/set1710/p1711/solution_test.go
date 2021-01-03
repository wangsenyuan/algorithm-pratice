package p1711

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := waysToSplit(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 2, 2, 5, 0}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 2, 1}
	expect := 0
	runSample(t, nums, expect)
}
