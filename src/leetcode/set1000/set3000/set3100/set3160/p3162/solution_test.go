package p3162

import "testing"

func runSample(t *testing.T, nums []int, queries [][]int, expect int) {
	res := maximumSumSubsequence(nums, queries)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 5, 9}
	queries := [][]int{
		{1, -2}, {0, -3},
	}
	expect := 21
	runSample(t, nums, queries, expect)
}
func TestSample2(t *testing.T) {
	nums := []int{0, -1}
	queries := [][]int{
		{0, -5},
	}
	expect := 0
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 0, -1, -2, 3, 1, -1}
	queries := [][]int{
		{3, 1}, {0, -2}, {1, -1}, {0, -2}, {5, 4}, {6, -3}, {6, -2}, {2, -1},
	}
	expect := 36
	runSample(t, nums, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{-4, 0, 1, -3, 0, -1, 1, -2, 2, -1}
	queries := [][]int{
		{5, -1}, {8, 3}, {4, 2}, {2, -4},
	}
	expect := 22
	runSample(t, nums, queries, expect)
}
