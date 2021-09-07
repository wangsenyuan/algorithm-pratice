package p1994

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := numberOfGoodSubsets(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 2, 3, 15}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{18, 28, 2, 17, 29, 30, 15, 9, 12}
	expect := 19
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{5, 10, 1, 26, 24, 21, 24, 23, 11, 12, 27, 4, 17, 16, 2, 6, 1, 1, 6, 8, 13, 30, 24, 20, 2, 19}
	expect := 5368
	runSample(t, nums, expect)
}
