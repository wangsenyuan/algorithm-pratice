package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 3},
		{3, 2},
	}
	expect := 7
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 4},
		{2, 3},
		{3, 1},
	}
	expect := 12
	runSample(t, n, edges, expect)
}
