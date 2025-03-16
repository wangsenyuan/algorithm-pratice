package p3489

import (
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect int) {
	res := minZeroArray(nums, queries)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 0, 2}
	queries := [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}
	expect := 2
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 2, 1}
	queries := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 4, 1}}
	expect := 4
	runSample(t, nums, queries, expect)
}
