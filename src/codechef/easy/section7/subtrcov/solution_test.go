package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, k int, expect int) {
	res := solve(n, E, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{5, 6},
		{3, 7},
	}
	expect := 2
	runSample(t, n, E, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 4
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
	}
	expect := 3
	runSample(t, n, E, k, expect)
}
