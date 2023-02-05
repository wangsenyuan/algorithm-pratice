package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect string) {
	res := solve(n, edges)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{3, 5},
		{1, 3},
		{3, 2},
		{4, 2},
	}
	expect := "01111"
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	edges := [][]int{
		{1, 2},
		{2, 5},
		{3, 4},
		{4, 2},
		{8, 1},
		{4, 5},
		{10, 5},
		{9, 5},
		{8, 2},
		{5, 7},
		{4, 6},
	}
	expect := "0011111011"
	runSample(t, n, edges, expect)
}
