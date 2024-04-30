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
		{1, 3},
		{2, 6},
		{0, 4},
		{3, 3},
	}
	expect := 1
	runSample(t, segs, expect)
}
