package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect int) {
	res := solve(n, roads)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	roads := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{3, 3, 3},
		{2, 1, 4},
		{2, 2, 5},
	}
	expect := 4
	runSample(t, n, roads, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	roads := [][]int{
		{1, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{2, 4, 1},
		{3, 5, 1},
		{4, 5, 1},
		{5, 5, 1},
	}
	expect := -1
	runSample(t, n, roads, expect)
}
