package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, E [][]int, expect int) {
	res := solve(n, A, E, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 6, 2
	A := []int{1, 4, 2, 2, 1, 5}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	expect := 6
	runSample(t, n, k, A, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 3
	A := []int{1, 13, 5, 14, 2, 9, 0, 8}
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{3, 6},
		{3, 7},
		{6, 8},
	}
	expect := 4
	runSample(t, n, k, A, E, expect)
}
