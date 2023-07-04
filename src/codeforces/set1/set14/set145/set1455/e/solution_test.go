package main

import "testing"

func runSample(t *testing.T, pts [][]int, expect int) {
	res := int(solve(pts))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pts := [][]int{
		{0, 2},
		{4, 2},
		{2, 0},
		{2, 4},
	}
	expect := 8
	runSample(t, pts, expect)
}

func TestSample2(t *testing.T) {
	pts := [][]int{
		{1, 0},
		{2, 0},
		{4, 0},
		{6, 0},
	}
	expect := 7
	runSample(t, pts, expect)
}

func TestSample3(t *testing.T) {
	pts := [][]int{
		{1, 6},
		{2, 2},
		{2, 5},
		{4, 1},
	}
	expect := 5
	runSample(t, pts, expect)
}
