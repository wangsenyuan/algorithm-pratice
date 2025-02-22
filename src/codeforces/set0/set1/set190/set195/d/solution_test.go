package main

import "testing"

func runSample(t *testing.T, fns [][]int, expect int) {
	res := solve(fns)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	fns := [][]int{
		{1, 0},
	}
	expect := 1
	runSample(t, fns, expect)
}

func TestSample2(t *testing.T) {
	fns := [][]int{
		{1, 0},
		{0, 2},
		{-1, 1},
	}
	expect := 2
	runSample(t, fns, expect)
}

func TestSample3(t *testing.T) {
	fns := [][]int{
		{-2, -4},
		{1, 7},
		{-5, 1},
	}
	expect := 3
	runSample(t, fns, expect)
}
