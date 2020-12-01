package main

import "testing"

func runSample(t *testing.T, n, m, k int, E [][]int, expect int) {
	res := solve(n, m, k, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 4, 6, 3
	E := [][]int{
		{4, 2, 1},
		{1, 2, 2},
		{2, 4, 3},
		{4, 1, 4},
		{4, 3, 5},
		{3, 1, 6},
	}
	expect := 2
	runSample(t, n, m, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 5, 5, 1
	E := [][]int{
		{1, 4, 1},
		{5, 1, 2},
		{2, 5, 3},
		{4, 3, 4},
		{3, 2, 5},
	}
	expect := 1
	runSample(t, n, m, k, E, expect)
}
