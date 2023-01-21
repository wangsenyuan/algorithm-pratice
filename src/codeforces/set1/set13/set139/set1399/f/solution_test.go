package main

import "testing"

func runSample(t *testing.T, n int, segs [][]int, expect int) {
	res := solve(n, segs)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, segs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	segs := [][]int{
		{1, 5},
		{2, 4},
		{2, 3},
		{3, 4},
	}
	expect := 3
	runSample(t, n, segs, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	segs := [][]int{
		{1, 5},
		{2, 3},
		{2, 5},
		{3, 5},
		{2, 2},
	}
	expect := 4
	runSample(t, n, segs, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	segs := [][]int{
		{1, 3},
		{2, 4},
		{2, 3},
	}
	expect := 2
	runSample(t, n, segs, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	segs := [][]int{
		{1, 10},
		{2, 8},
		{2, 5},
		{3, 4},
		{4, 4},
		{6, 8},
		{7, 7},
	}
	expect := 7
	runSample(t, n, segs, expect)
}
