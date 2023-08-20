package p2831

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := longestEqualSubarray(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2, 3, 1, 3}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 2, 2, 1, 1}
	k := 2
	expect := 4
	runSample(t, nums, k, expect)
}
