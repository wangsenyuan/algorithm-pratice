package p2419

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestSubarray(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 3, 2, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{100, 5, 5}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 279979}
	expect := 1
	runSample(t, nums, expect)
}
