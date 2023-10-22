package p2909

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minGroupsForValidAssignment(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1, 3, 1, 2, 2, 2, 3}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 1, 2, 2, 3, 1, 3, 1, 1, 1, 1, 2}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 10, 10, 3, 1, 1}
	expect := 4
	runSample(t, nums, expect)
}
