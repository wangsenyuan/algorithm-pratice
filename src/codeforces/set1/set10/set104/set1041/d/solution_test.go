package main

import "testing"

func runSample(t *testing.T, h int, segs [][]int, expect int) {
	res := solve(h, segs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 4
	segs := [][]int{
		{2, 5},
		{7, 9},
		{10, 11},
	}
	expect := 10
	runSample(t, h, segs, expect)
}

func TestSample2(t *testing.T) {
	h := 10
	segs := [][]int{
		{5, 7},
		{11, 12},
		{16, 20},
		{25, 26},
		{30, 33},
	}
	expect := 18
	runSample(t, h, segs, expect)
}

func TestSample3(t *testing.T) {
	h := 1000000000
	segs := [][]int{
		{1, 1000000000},
	}
	expect := 1999999999
	runSample(t, h, segs, expect)
}

func TestSample4(t *testing.T) {
	h := 10
	segs := [][]int{
		{3, 4},
		{10, 11},
		{12, 13},
		{14, 16},
		{18, 20},
	}
	expect := 16
	runSample(t, h, segs, expect)
}
