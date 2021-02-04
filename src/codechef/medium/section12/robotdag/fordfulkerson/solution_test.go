package main

import "testing"

func runSample(t *testing.T, n, m, k int, E [][]int, expect int) {
	res := solve(n, m, k, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 8, 11, 1
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{6, 7},
		{2, 5},
		{3, 6},
		{3, 2},
		{4, 6},
		{5, 7},
		{7, 8},
		{2, 7},
	}
	expect := 3

	runSample(t, n, m, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 8, 11, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{6, 7},
		{2, 5},
		{3, 6},
		{3, 2},
		{4, 6},
		{5, 7},
		{7, 8},
		{2, 7},
	}
	expect := 4

	runSample(t, n, m, k, E, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 8, 11, 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{6, 7},
		{2, 5},
		{3, 6},
		{3, 2},
		{4, 6},
		{5, 7},
		{7, 8},
		{2, 7},
	}
	expect := 5

	runSample(t, n, m, k, E, expect)
}

