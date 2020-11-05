package main

import "testing"

func runSample(t *testing.T, n int, detachments [][]int, expect int) {
	res := solve(n, detachments)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	ss := [][]int{
		{100, 0},
		{0, 100},
		{-100, 0},
		{0, -100},
	}
	expect := 100
	runSample(t, n, ss, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	ss := [][]int{
		{0, 2},
		{1, 0},
		{-3, 0},
		{0, -2},
		{-1, -1},
		{-1, -3},
		{-2, -3},
	}
	expect := -1
	runSample(t, n, ss, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	ss := [][]int{
		{0, 0},
		{0, -1},
		{3, 0},
		{-2, 0},
		{-2, 1},
	}
	expect := 2
	runSample(t, n, ss, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	ss := [][]int{
		{0, 0},
		{2, 0},
		{0, -1},
		{-2, 0},
		{-2, 1},
	}
	expect := 2
	runSample(t, n, ss, expect)
}
