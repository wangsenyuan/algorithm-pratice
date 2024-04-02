package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	_, res := solve(nums)

	a := countDiff(nums, expect)
	b := countDiff(nums, res)

	if a != b {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func countDiff(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	var res int
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	a := []int{-10, -9, -10, -8, -10, -9, -9}
	expect := []int{-10, -10, -9, -9, -9, -9, -9}
	runSample(t, a, expect)
}
