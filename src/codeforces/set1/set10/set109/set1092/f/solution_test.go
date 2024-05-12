package main

import "testing"

func runSample(t *testing.T, n int, a []int, edges [][]int, expect int) {
	res := solve(n, a, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	a := []int{9, 4, 1, 7, 10, 1, 6, 5}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{1, 5},
		{5, 6},
		{5, 7},
		{5, 8},
	}
	expect := 121
	runSample(t, n, a, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	a := []int{1337}

	expect := 0
	runSample(t, n, a, nil, expect)
}
