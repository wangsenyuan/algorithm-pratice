package main

import "testing"

func runSample(t *testing.T, n int, neighbors [][]int, expect int) {
	res := solve(n, neighbors)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := [][]int{
		{1, 2, 3},
		{2, 3, 1},
	}
	expect := 4
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	a := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 1, 4, 5},
		{3, 4, 5, 1, 2},
		{3, 5, 4, 2, 1},
		{2, 3, 5, 4, 1},
		{1, 2, 3, 4, 5},
	}
	expect := 5
	runSample(t, n, a, expect)
}
