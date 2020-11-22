package p5607

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := waysToMakeFair(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 6, 4}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 0
	runSample(t, nums, expect)
}
