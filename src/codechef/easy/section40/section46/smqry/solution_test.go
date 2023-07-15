package main

import "testing"

func runSample(t *testing.T, n int, Q [][]int, expect int) {
	res := solve(n, Q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	Q := [][]int{
		{1, 2},
		{3, 2},
	}
	expect := 3
	runSample(t, n, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	Q := [][]int{
		{1, 2},
		{3, 2},
		{4 ^ 3, 1 ^ 3},
	}
	expect := 0
	runSample(t, n, Q, expect)
}
