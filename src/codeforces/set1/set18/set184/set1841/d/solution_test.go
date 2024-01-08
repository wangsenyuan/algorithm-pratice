package main

import "testing"

func runSample(t *testing.T, segs [][]int, expect int) {
	res := solve(segs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{2, 4},
		{9, 12},
		{2, 4},
		{7, 7},
		{4, 8},
		{10, 13},
		{6, 8},
	}
	expect := 1
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{2, 2},
		{2, 8},
		{0, 10},
		{1, 2},
		{5, 6},
	}
	expect := 3
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}
	expect := 4
	runSample(t, segs, expect)
}
