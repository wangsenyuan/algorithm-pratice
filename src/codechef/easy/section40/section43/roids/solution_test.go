package main

import "testing"

func runSample(t *testing.T, roids [][]int, expect int) {
	res := solve(roids)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	roids := [][]int{
		{1, 2},
		{3, 1},
		{3, 4},
		{3, 7},
		{4, 6},
		{4, 2},
	}
	expect := 616666671
	runSample(t, roids, expect)
}
