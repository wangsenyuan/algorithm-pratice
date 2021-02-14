package p1760

import "testing"

func runSample(t *testing.T, nums []int, operations int, expect int) {
	res := minimumSize(nums, operations)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{9}
	operations := 2
	expect := 3
	runSample(t, nums, operations, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 4, 8, 2}
	operations := 4
	expect := 2
	runSample(t, nums, operations, expect)
}
