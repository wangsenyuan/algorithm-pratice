package main

import "testing"

func runSample(t *testing.T, n, k int, segs [][]int, expect int) {
	res := solve(n, k, segs)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, segs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	segs := [][]int{
		{1, 6},
		{2, 4},
		{3, 6},
	}
	expect := 3
	runSample(t, n, k, segs, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	segs := [][]int{
		{1, 6},
		{2, 3},
		{3, 6},
	}
	expect := 3
	runSample(t, n, k, segs, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 2
	segs := [][]int{
		{1, 5},
		{2, 3},
		{3, 6},
	}
	expect := 2
	runSample(t, n, k, segs, expect)
}

func TestSample4(t *testing.T) {
	n, k := 3, 2
	segs := [][]int{
		{1, 2},
		{2, 3},
		{3, 6},
	}
	expect := 0
	runSample(t, n, k, segs, expect)
}

func TestSample5(t *testing.T) {
	n, k := 3, 2
	segs := [][]int{
		{1, 5},
		{3, 6},
		{2, 3},
	}
	expect := 2
	runSample(t, n, k, segs, expect)
}
