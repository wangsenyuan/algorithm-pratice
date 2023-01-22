package p2547

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minCost(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1, 3, 3}
	k := 2
	expect := 8
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1}
	k := 2
	expect := 6
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1}
	k := 5
	expect := 10
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 3, 3, 3, 4, 5, 4, 6, 2, 4, 2, 1, 5, 6, 4, 5, 1, 1, 3, 3}
	k := 1
	expect := 9
	runSample(t, nums, k, expect)
}
