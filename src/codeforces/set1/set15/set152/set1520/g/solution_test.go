package main

import "testing"

func runSample(t *testing.T, w int, a [][]int, expect int) {
	res := solve(w, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var w = 1
	a := [][]int{
		{0, -1, 0, 1, -1},
		{0, 20, 0, 0, -1},
		{-1, -1, -1, -1, -1},
		{3, 0, 0, 0, 0},
		{-1, 0, 0, 0, 0},
	}
	var expect = 14
	runSample(t, w, a, expect)
}

func TestSample2(t *testing.T) {
	var w = 2
	a := [][]int{
		{0, -1},
		{0, 10},
		{8, 4},
		{0, 0},
		{0, 0},
		{-1, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{4, 0},
	}
	var expect = 16
	runSample(t, w, a, expect)
}

func TestSample3(t *testing.T) {
	var w = 2
	a := [][]int{
		{0, 0, -1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, -1, 0, 0, 0},
		{-1, -1, 0, 0, 0, 0, 0},
		{0, -1, 0, 0, -1, 0, 0},
		{-1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{-1, 0, 0, 0, 0, -1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	var expect = 30
	runSample(t, w, a, expect)
}
