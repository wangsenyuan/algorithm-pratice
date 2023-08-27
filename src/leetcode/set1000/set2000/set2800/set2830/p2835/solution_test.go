package p2835

import "testing"

func runSample(t *testing.T, nums []int, target int, expect int) {
	res := minOperations(nums, target)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 8}
	target := 7
	expect := 1
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 32, 1, 2}
	target := 12
	expect := 2
	runSample(t, nums, target, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 32, 1}
	target := 35
	expect := -1
	runSample(t, nums, target, expect)
}
