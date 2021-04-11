package p154

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := findMin(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 2, 2, 0, 1}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 0, 1, 1, 1}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 1, 3, 3}
	expect := 1
	runSample(t, nums, expect)
}
