package main

import "testing"

func runSample(t *testing.T, n, m, a, b, c int, P []int, edges [][]int, expect int64) {
	res := solve(n, m, a, b, c, P, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, a, b, c := 4, 3, 2, 3, 4
	P := []int{1, 2, 3}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := int64(7)
	runSample(t, n, m, a, b, c, P, edges, expect)
}

func TestSample2(t *testing.T) {
	n, m, a, b, c := 7, 9, 1, 5, 7
	P := []int{2, 10, 4, 8, 5, 6, 7, 3, 3}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{3, 2},
		{3, 5},
		{4, 2},
		{5, 6},
		{1, 7},
		{6, 7},
	}
	expect := int64(12)
	runSample(t, n, m, a, b, c, P, edges, expect)
}
