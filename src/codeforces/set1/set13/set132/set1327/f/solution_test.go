package main

import "testing"

func runSample(t *testing.T, n int, k int, Q [][]int, expect int) {
	res := solve(n, k, Q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	Q := [][]int{
		{1, 3, 3},
		{3, 4, 6},
	}
	expect := 3
	runSample(t, n, k, Q, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 2
	Q := [][]int{
		{1, 3, 2},
		{2, 5, 0},
		{3, 3, 3},
	}
	expect := 33
	runSample(t, n, k, Q, expect)
}
