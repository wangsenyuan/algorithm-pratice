package p2787

import "testing"

func runSample(t *testing.T, nums []int, x int, expect int) {
	res := int(maxScore(nums, x))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 50, 65, 85, 8, 73, 55, 50, 29, 95, 5, 68, 52, 79}
	x := 74
	expect := 470
	runSample(t, nums, x, expect)
}
