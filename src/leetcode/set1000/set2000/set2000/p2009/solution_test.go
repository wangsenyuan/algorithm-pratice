package p2009

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minOperations(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 5, 6}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 2, 5, 3}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 10, 100, 1000}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{8, 5, 9, 9, 8, 4}
	expect := 2
	runSample(t, nums, expect)
}
