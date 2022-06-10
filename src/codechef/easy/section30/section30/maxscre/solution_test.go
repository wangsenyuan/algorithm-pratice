package main

import "testing"

func runSample(t *testing.T, n int, S [][]int, expect int) {
	res := int(solve(n, S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := [][]int{
		{2, 2, 5},
		{2, 2, 10},
		{3, 3, 14},
		{3, 3, 20},
	}
	expect := 30
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := [][]int{
		{1, 3, 4},
		{2, 3, 5},
		{1, 1, 10},
		{1, 4, 13},
	}
	expect := 32
	runSample(t, n, S, expect)
}
