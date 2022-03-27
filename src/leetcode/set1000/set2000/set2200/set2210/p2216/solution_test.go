package p2216

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minDeletion(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 3, 5}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	expect := 2
	runSample(t, nums, expect)
}
