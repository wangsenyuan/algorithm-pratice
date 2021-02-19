package main

import "testing"

func runSample(t *testing.T, N, Q, M int, A []int, queries [][]int, expect int) {
	res := solve(N, Q, M, A, queries)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q, m := 3, 2, 10
	A := []int{2, 3, 4}

	queries := [][]int{
		{1, 2}, {2, 3},
	}
	expect := 2
	runSample(t, n, q, m, A, queries, expect)
}
