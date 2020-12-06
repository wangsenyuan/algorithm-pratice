package p5618

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxOperations(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 3, 4, 3}
	k := 6
	expect := 1
	runSample(t, nums, k, expect)
}
