package p2734

import "testing"

func runSample(t *testing.T, nums []int, x int, expect int64) {
	res := minCost(nums, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{20, 1, 15}
	x := 5
	var expect int64 = 13
	runSample(t, nums, x, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3}
	x := 4
	var expect int64 = 6
	runSample(t, nums, x, expect)
}
