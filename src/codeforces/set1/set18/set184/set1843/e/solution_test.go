package main

import "testing"

func runSample(t *testing.T, n int, segs [][]int, changes []int, expect int) {
	res := solve(n, segs, changes)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	segs := [][]int{
		{1, 2},
		{4, 5},
		{1, 5},
		{1, 3},
		{2, 4},
	}
	changes := []int{5, 3, 1, 2, 4}
	expect := 3
	runSample(t, n, segs, changes, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	segs := [][]int{
		{1, 1},
		{4, 4},
	}
	changes := []int{2, 3}
	expect := -1
	runSample(t, n, segs, changes, expect)
}
