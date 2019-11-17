package p1262

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxSumDivThree(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 6, 5, 1, 8}
	expect := 18
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4}
	expect := 12
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 6, 2, 2, 7}
	expect := 15
	runSample(t, nums, expect)
}
