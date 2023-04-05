package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, Q [][]int, expect int) {
	res := solve(n, m, k, Q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 1, 1, 3
	Q := [][]int{
		{1, 1},
		{1, 1},
	}
	expect := k
	runSample(t, n, m, k, Q, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 2, 2, 2
	Q := [][]int{
		{2, 1},
		{1, 1},
		{2, 2},
	}
	expect := 4
	runSample(t, n, m, k, Q, expect)
}
