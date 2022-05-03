package main

import "testing"

func runSample(t *testing.T, n int, rects [][]int, expect int) {
	res := solve(n, rects)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	rects := [][]int{
		{0, 5, 1, 2},
		{1, 2, 0, 3},
		{3, 4, 0, 3},
	}
	expect := 2
	runSample(t, n, rects, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	rects := [][]int{
		{-3, 3, -1, 1},
		{-2, 2, -2, 2},
		{-1, 1, -3, 3},
	}
	expect := 1
	runSample(t, n, rects, expect)
}
