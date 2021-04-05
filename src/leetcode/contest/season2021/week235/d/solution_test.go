package d

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countDifferentSubsequenceGCDs(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 15, 40, 5, 6}
	expect := 7
	runSample(t, nums, expect)
}
