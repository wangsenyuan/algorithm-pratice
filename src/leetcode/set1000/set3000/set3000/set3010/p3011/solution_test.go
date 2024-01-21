package p3011

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minimumArrayLength(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 5, 5, 10, 5}
	expect := 2
	runSample(t, nums, expect)
}
