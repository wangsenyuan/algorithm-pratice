package main

import "testing"

func runSample(t *testing.T, n int, pairs [][]int, expect int) {
	res := solve(n, pairs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	pairs := [][]int{
		{1, 3},
		{2, 3},
	}
	expect := 4
	runSample(t, n, pairs, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	pairs := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 5
	runSample(t, n, pairs, expect)
}
