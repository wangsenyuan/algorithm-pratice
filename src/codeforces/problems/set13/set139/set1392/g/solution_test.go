package main

import "testing"

func runSample(tc *testing.T, n, m, k int, a, b string, swaps [][]int, expect int) {
	res := solve(n, m, k, a, b, swaps)

	if res[0] != expect {
		tc.Fatalf("Sample expect %d, but got %d", expect, res[0])
	}
	if res[2]-res[1]+1 < m {
		tc.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 4, 2, 5
	a, b := "11000", "00011"
	swaps := [][]int{
		{1, 3},
		{3, 5},
		{4, 2},
		{3, 4},
	}
	expect := 5
	runSample(t, n, m, k, a, b, swaps, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 4, 3, 5
	a, b := "11000", "00011"
	swaps := [][]int{
		{1, 3},
		{1, 5},
		{2, 4},
		{1, 5},
	}
	expect := 3
	runSample(t, n, m, k, a, b, swaps, expect)
}
