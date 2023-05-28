package p2707

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := canTraverseAllPairs(nums)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 9, 5}
	expect := false
	runSample(t, nums, expect)
}
