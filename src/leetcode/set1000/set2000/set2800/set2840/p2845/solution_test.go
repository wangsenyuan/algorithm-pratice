package p2845

import "testing"

func runSample(t *testing.T, nums []int, mod int, k int, expect int) {
	res := countInterestingSubarrays(nums, mod, k)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{11, 12, 21, 31}
	mod := 10
	k := 1
	expect := 5
	runSample(t, nums, mod, k, expect)
}
