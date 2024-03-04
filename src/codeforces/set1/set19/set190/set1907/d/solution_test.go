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
		{1, 5},
		{3, 4},
		{5, 6},
		{8, 10},
		{0, 1},
	}
	expect := 7
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{0, 2},
		{0, 1},
		{0, 3},
	}
	expect := 0
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{3, 8},
		{10, 18},
		{6, 11},
	}
	expect := 5
	runSample(t, segs, expect)
}

func TestSample4(t *testing.T) {
	segs := [][]int{
		{10, 20},
		{0, 5},
		{15, 17},
		{2, 2},
	}
	expect := 13
	runSample(t, segs, expect)
}
