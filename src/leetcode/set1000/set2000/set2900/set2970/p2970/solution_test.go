package p2970

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := incremovableSubarrayCount(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 10
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 5, 7, 8}
	expect := 7
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{8, 7, 6, 6}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3, 1, 2, 3, 4, 5, 6}
	expect := 25
	runSample(t, nums, expect)
}
