package p2025

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := waysToPartition(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, -1, 2}
	k := 3
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 0, 0}
	k := 0
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{22, 4, -25, -20, -15, 15, -16, 7, 19, -10, 0, -13, -14}
	k := -33
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -93166, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 93166, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 69428
	expect := 37
	runSample(t, nums, k, expect)
}
