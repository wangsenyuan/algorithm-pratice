package p2172

import "testing"

func runSample(t *testing.T, nums []int, numSlots int, expect int) {
	res := maximumANDSum(nums, numSlots)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	numSlots := 3
	expect := 9
	runSample(t, nums, numSlots, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 10, 4, 7, 1}
	numSlots := 9
	expect := 24
	runSample(t, nums, numSlots, expect)
}
