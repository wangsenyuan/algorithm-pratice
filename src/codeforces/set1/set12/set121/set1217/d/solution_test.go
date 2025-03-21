package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	cnt, res := solve(n, edges)
	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, cnt, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{2, 4},
		{1, 4},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	expect := 2
	runSample(t, n, edges, expect)
}
