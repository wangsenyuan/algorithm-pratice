package p1567

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := getMaxLen(nums)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -2, -3, 4}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 1, -2, -3, -4}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-1, -2, -3, 0, 1}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{50, 4, 36, 12, 32, -49, 62, 22, 53, -67, -12, 49, 55, 0, 9, 76, -26, 10, 35, 16, -11, -10, 92, -32, -72, 0, 73, -91, 52, 52, -3, 67, -37, 26, 89, -33, 0, -46, 0, -16, 8, 61, 0, 79, 20, 48}
	expect := 10
	runSample(t, nums, expect)
}
