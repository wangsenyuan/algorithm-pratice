package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res, ops := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	if len(ops) != n-1 {
		t.Fatalf("Sample result %v, not correct", ops)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 9
	runSample(t, n, edges, expect)
}
