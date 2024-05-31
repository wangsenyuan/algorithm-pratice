package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, s []int, expect int) {
	res := solve(n, edges, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 2},
		{3, 2, 1},
		{2, 4, 5},
		{2, 5, 7},
		{4, 5, 1},
	}
	s := []int{5, 2, 1, 3, 3}
	expect := 19
	runSample(t, n, edges, s, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2, 5},
		{1, 3, 5},
		{1, 4, 4},
		{1, 5, 8},
		{2, 3, 6},
		{2, 4, 3},
		{2, 5, 2},
		{3, 4, 1},
		{3, 5, 8},
		{4, 5, 2},
	}
	s := []int{7, 2, 8, 4, 1}
	expect := 36
	runSample(t, n, edges, s, expect)
}
