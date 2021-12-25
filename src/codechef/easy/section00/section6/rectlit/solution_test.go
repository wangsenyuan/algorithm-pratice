package main

import "testing"

func runSample(t *testing.T, k, n int, coords [][]int, expect bool) {
	res := solve(k, n, coords)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", k, n, coords, expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, n := 2, 10
	coords := [][]int{
		{0, 0},
		{1, 0},
	}
	expect := true
	runSample(t, k, n, coords, expect)
}

func TestSample2(t *testing.T) {
	k, n := 2, 10
	coords := [][]int{
		{1, 2},
		{1, 1},
	}
	expect := false
	runSample(t, k, n, coords, expect)
}
