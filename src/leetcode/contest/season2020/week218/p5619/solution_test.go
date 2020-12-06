package p5619

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minimumIncompatibility(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 4}
	k := 2
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 3, 8, 1, 3, 1, 2, 2}
	k := 4
	expect := 6
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 1, 2, 5, 6, 1}
	k := 2
	expect := 8
	runSample(t, nums, k, expect)
}
