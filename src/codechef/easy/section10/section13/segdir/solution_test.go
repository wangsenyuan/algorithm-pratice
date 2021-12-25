package main

import "testing"

func runSample(t *testing.T, n int, S [][]int, expect bool) {
	res := solve(n, S)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := [][]int{
		{5, 7, 2},
		{4, 6, 1},
		{1, 5, 2},
		{6, 8, 1},
	}
	expect := true
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	S := [][]int{
		{1, 2, 1},
		{1, 2, 1},
		{1, 2, 1},
		{1, 2, 1},
	}
	expect := false
	runSample(t, n, S, expect)
}
