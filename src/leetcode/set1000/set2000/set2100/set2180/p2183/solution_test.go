package p2183

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := coutPairs(nums, k)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	expect := 7
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := 5
	expect := 0
	runSample(t, nums, k, expect)
}
