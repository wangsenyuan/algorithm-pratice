package p2962

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := countSubarrays(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2, 3, 3}
	k := 2
	expect := 6
	runSample(t, nums, k, expect)
}
