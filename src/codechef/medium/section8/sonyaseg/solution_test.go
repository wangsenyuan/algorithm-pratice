package main

import "testing"

func runSample(t *testing.T, n, k int, seg [][]int, expect int) {
	res := solve(n, k, seg)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, seg, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	segs := [][]int{
		{1, 4},
		{2, 5},
		{5, 10},
		{4, 7},
		{8, 9},
	}
	expect := 4
	runSample(t, n, k, segs, expect)
}

func TestSample2(t *testing.T) {
	n, k := 20, 5
	segs := [][]int{
		{6, 7},
		{4, 12},
		{3, 7},
		{2, 6},
		{3, 4},
		{2, 13},
		{5, 8},
		{8, 11},
		{11, 14},
		{4, 12},
		{8, 14},
		{12, 14},
		{12, 15},
		{2, 3},
		{6, 8},
		{8, 13},
		{4, 4},
		{11, 13},
		{11, 12},
		{4, 8},
	}
	expect := 14891
	runSample(t, n, k, segs, expect)
}
