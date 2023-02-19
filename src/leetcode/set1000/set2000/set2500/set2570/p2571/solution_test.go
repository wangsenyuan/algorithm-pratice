package p2571

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := squareFreeSubsets(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, 4, 5}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{11, 2, 19, 7, 9, 27}
	expect := 15
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{17, 27, 20, 1, 19}
	expect := 7
	runSample(t, nums, expect)
}
