package main

import "testing"

func runSample(t *testing.T, n int, C [][]int, expect int) {
	res := solve(n, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	C := [][]int{
		{2, 0},
		{0, 2},
	}
	expect := 3
	runSample(t, n, C, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	C := [][]int{
		{1, 0},
		{2, 1},
		{0, 1},
		{1, 2},
	}
	expect := 6
	runSample(t, n, C, expect)
}
