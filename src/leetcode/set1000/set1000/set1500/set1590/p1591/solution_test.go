package p1591

import "testing"

func runSample(t *testing.T, nums []int, requests [][]int, expect int) {
	res := maxSumRangeQuery(nums, requests)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	requests := [][]int{{1, 3}, {0, 1}}
	expect := 19
	runSample(t, nums, requests, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	requests := [][]int{{0, 1}}
	expect := 11
	runSample(t, nums, requests, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 10}
	requests := [][]int{{0, 2}, {1, 3}, {1, 1}}
	expect := 47
	runSample(t, nums, requests, expect)
}
