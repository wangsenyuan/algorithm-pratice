package p2856

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minLengthAfterRemovals(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 4, 9}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 4, 4}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4, 4}
	expect := 0
	runSample(t, nums, expect)
}
