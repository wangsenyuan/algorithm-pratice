package p1856

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxSumMinProduct(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	expect := 14
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 5, 6, 4, 2}
	expect := 60
	runSample(t, nums, expect)
}
