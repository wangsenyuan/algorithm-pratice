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
		{1, 4},
		{2, 3},
		{3, 6},
	}
	expect := 0
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 5},
	}
	expect := 1
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 2},
		{3, 8},
		{4, 5},
		{6, 7},
		{9, 10},
	}
	expect := 2
	runSample(t, segs, expect)
}

func TestSample4(t *testing.T) {
	segs := [][]int{
		{1, 5},
		{2, 4},
		{3, 5},
		{3, 8},
		{4, 8},
	}
	expect := 0
	runSample(t, segs, expect)
}
