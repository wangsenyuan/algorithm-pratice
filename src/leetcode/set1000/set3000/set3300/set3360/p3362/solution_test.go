package p3362

import "testing"

func runSample(t *testing.T, nums []int, queries [][]int, expect int) {
	res := maxRemoval(nums, queries)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 0, 2}
	queries := [][]int{
		{0, 2},
		{0, 2},
		{1, 1},
	}
	expect := 1
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	queries := [][]int{
		{0, 3},
	}
	expect := -1
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{0, 0, 1, 1, 0}
	queries := [][]int{
		{3, 4},
		{0, 2},
		{2, 3},
	}
	expect := 2
	runSample(t, nums, queries, expect)
}
