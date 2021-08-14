package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{3, 1},
		{3, 4},
		{3, 5},
		{2, 6},
		{2, 7},
	}
	expect := 8
	runSample(t, n, E, expect)
}
