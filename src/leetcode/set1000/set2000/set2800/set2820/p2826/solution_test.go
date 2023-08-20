package p2826

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minimumOperations(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3, 2, 1}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 2, 1, 3, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 2, 2, 3, 3}
	expect := 0
	runSample(t, nums, expect)
}
