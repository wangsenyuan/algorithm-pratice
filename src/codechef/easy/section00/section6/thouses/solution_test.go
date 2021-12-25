package main

import "testing"

func runSample(t *testing.T, n int, x int, E [][]int, expect int) {
	res := solve(n, x, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 4, 1
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 7
	runSample(t, n, x, E, expect)
}

func TestSample2(t *testing.T) {
	n, x := 8, 1
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{5, 6},
		{5, 7},
		{7, 8},
	}
	expect := 11
	runSample(t, n, x, E, expect)
}
