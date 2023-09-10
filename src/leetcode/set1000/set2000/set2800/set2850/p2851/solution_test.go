package p2851

import "testing"

func runSample(t *testing.T, nums [][]int, expect int) {
	res := numberOfPoints(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := [][]int{
		{4, 4}, {9, 10}, {9, 10}, {3, 8},
	}
	expect := 8
	runSample(t, nums, expect)
}
