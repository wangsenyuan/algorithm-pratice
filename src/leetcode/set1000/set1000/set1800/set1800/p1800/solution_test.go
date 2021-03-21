package p1800

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxAscendingSum(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 20, 30, 5, 10, 50}
	expect := 65
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{100, 10, 1}
	expect := 100
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{12, 17, 15, 13, 10, 11, 12}
	expect := 33
	runSample(t, nums, expect)
}
