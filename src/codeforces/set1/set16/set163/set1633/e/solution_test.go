package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, k int, a, b, c int, Q []int, expect int64) {
	res := solve(n, E, k, a, b, c, Q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{4, 1, 4},
		{3, 1, 0},
		{3, 5, 3},
		{2, 5, 4},
		{3, 4, 8},
		{4, 3, 4},
		{4, 2, 8},
		{5, 3, 9},
	}
	k, a, b, c := 11, 1, 1, 10
	Q := []int{0, 1, 2}
	var expect int64 = 4
	runSample(t, n, E, k, a, b, c, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2, 50},
		{2, 3, 100},
		{1, 3, 150},
	}
	k, a, b, c := 10000000, 0, 0, 100000000
	Q := []int{75}
	var expect int64 = 164
	runSample(t, n, E, k, a, b, c, Q, expect)
}
