package main

import "testing"

func runSample(t *testing.T, n int, nodes [][]int, expect int) {
	res := solve(n, nodes)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	nodes := [][]int{
		{-1, 3},
		{1, 2},
		{1, 1},
		{1, 4},
		{4, 5},
		{4, 3},
		{5, 2},
	}
	expect := 17
	runSample(t, n, nodes, expect)
}

func TestSample2(t *testing.T) {
	n := 20
	nodes := [][]int{
		{-1, 100},
		{1, 10},
		{2, 26},
		{2, 33},
		{3, 31},
		{2, 28},
		{1, 47},
		{6, 18},
		{6, 25},
		{9, 2},
		{4, 17},
		{6, 18},
		{6, 2},
		{6, 30},
		{13, 7},
		{5, 25},
		{7, 11},
		{11, 7},
		{17, 40},
		{12, 43},
	}
	expect := 355
	runSample(t, n, nodes, expect)
}
