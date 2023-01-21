package main

import "testing"

func runSample(t *testing.T, n int, C [][]int, expect int) {
	res := int(solve(n, C))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	C := [][]int{
		{0, 8},
		{1, 99},
	}
	expect := 100
	runSample(t, n, C, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	C := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{9, 9, 2, 2},
		{9, 9, 9, 9},
	}
	expect := 22
	runSample(t, n, C, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	C := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 2, 2, 0, 2, 2},
		{0, 0, 2, 0, 1, 6, 2, 1},
		{0, 2, 0, 0, 2, 4, 7, 4},
		{2, 0, 0, 0, 2, 0, 1, 6},
	}
	expect := 42
	runSample(t, n, C, expect)
}
