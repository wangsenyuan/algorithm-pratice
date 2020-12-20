package p1696

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxResult(nums, k)

	if res != expect {
		t.Errorf("Sample %v, %d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -1, -2, 4, -7, 3}
	k := 2
	expect := 7
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, -5, -2, 4, 0, 3}
	k := 3
	expect := 17
	runSample(t, nums, k, expect)
}
func TestSample3(t *testing.T) {
	nums := []int{1, -5, -20, 4, -1, 3, -6, -3}
	k := 2
	expect := 0
	runSample(t, nums, k, expect)
}
