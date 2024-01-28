package p3019

import "testing"

func runSample(t *testing.T, k int, nums []int, expect int) {
	res := minOrAfterOperations(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 5, 3, 2, 7}
	k := 2
	expect := 3
	runSample(t, k, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 3, 15, 14, 2, 8}
	k := 4
	expect := 2
	runSample(t, k, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 7, 10, 3, 9, 14, 9, 4}
	k := 1
	expect := 15
	runSample(t, k, nums, expect)
}
