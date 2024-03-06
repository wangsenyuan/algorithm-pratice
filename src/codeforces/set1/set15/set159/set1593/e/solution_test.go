package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, expect int) {
	res := solve(n, edges, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 14, 1
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{4, 6},
		{2, 7},
		{7, 8},
		{8, 9},
		{8, 10},
		{3, 11},
		{3, 12},
		{1, 13},
		{13, 14},
	}
	expect := 7
	runSample(t, n, edges, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 1000
	edges := [][]int{
		{1, 2},
	}
	expect := 0
	runSample(t, n, edges, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 2
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 0
	runSample(t, n, edges, k, expect)
}
