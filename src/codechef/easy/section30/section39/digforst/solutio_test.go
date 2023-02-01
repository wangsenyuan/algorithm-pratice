package main

import "testing"

func runSample(t *testing.T, n int, S [][]int, V []int, expect int) {
	res := solve(n, S, V)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	V := []int{1, 2, 1, 4}
	S := [][]int{
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 0},
	}
	expect := 7
	runSample(t, n, S, V, expect)
}
