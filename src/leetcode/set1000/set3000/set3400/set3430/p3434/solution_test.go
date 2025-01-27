package p3434

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxFrequency(nums, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", nums, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 1
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 2, 3, 4, 5, 5, 4, 3, 2, 2}
	k := 10
	expect := 4
	runSample(t, nums, k, expect)
}
