package p3385

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := beautifulSplits(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 1}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 0, 1, 3, 2, 2, 2}
	expect := 8
	runSample(t, nums, expect)
}
