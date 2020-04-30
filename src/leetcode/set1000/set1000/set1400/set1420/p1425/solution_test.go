package p1425

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := constrainedSubsetSum(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 2, -10, 5, 20}
	k := 2
	expect := 37
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-1, -2, -3}
	k := 1
	expect := -1
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, -2, -10, -5, 20}
	k := 2
	expect := 23
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{-5266, 4019, 7336, -3681, -5767}
	k := 2
	expect := 11355
	runSample(t, nums, k, expect)
}
