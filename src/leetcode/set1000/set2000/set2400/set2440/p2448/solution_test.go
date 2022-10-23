package p2448

import "testing"

func runSample(t *testing.T, nums []int, cost []int, expect int64) {
	res := minCost(nums, cost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 5, 2}
	cost := []int{2, 3, 1, 14}
	expect := 8
	runSample(t, nums, cost, int64(expect))
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	cost := []int{4, 2, 8, 1, 3}
	expect := 0
	runSample(t, nums, cost, int64(expect))
}
