package p2576

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxNumOfMarkedIndices(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 5, 2, 4}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{9, 2, 5, 4}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{7, 6, 8}
	expect := 0
	runSample(t, nums, expect)
}
