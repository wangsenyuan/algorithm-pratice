package main

import "testing"

func runSample(t *testing.T, mines [][]int, k int, expect int) {
	res := solve(mines, k)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mines := [][]int{
		{0, 0, 1},
		{0, 1, 4},
		{1, 0, 2},
		{1, 1, 3},
		{2, 2, 9},
	}
	k := 0
	expect := 2
	runSample(t, mines, k, expect)
}

func TestSample2(t *testing.T) {
	mines := [][]int{
		{0, 0, 1},
		{0, 1, 4},
		{1, 0, 2},
		{1, 1, 3},
		{2, 2, 9},
	}
	k := 2
	expect := 1
	runSample(t, mines, k, expect)
}

func TestSample3(t *testing.T) {
	mines := [][]int{
		{1, -1, 3},
		{0, -1, 9},
		{0, 1, 7},
		{-1, 0, 1},
		{-1, 1, 9},
		{-1, -1, 7},
	}
	k := 1
	expect := 0
	runSample(t, mines, k, expect)
}
