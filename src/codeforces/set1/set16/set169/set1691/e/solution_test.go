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
		{0, 0, 5},
		{1, 2, 12},
		{0, 4, 7},
		{1, 9, 16},
		{0, 13, 19},
	}
	expect := 2
	runSample(t, segs, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{1, 0, 1},
		{1, 1, 2},
		{0, 3, 4},
	}
	expect := 3
	runSample(t, segs, expect)
}

func TestSample3(t *testing.T) {
	segs := [][]int{
		{0, 82, 99},
		{1, 76, 98},
		{1, 14, 56},
		{1, 65, 72},
		{1, 81, 87},
		{0, 63, 71},
		{0, 58, 91},
		{0, 98, 99},
		{1, 44, 47},
		{1, 75, 84},
		{0, 56, 96},
		{0, 100, 100},
		{1, 87, 90},
		{1, 82, 84},
	}
	expect := 3
	runSample(t, segs, expect)
}
