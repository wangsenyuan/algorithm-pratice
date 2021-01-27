package main

import "testing"

func runSample(t *testing.T, n, k int, C []int, E [][]int, expect int) {
	res := solve(n, k, C, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	C := []int{1, 2, 0, 1, 2}
	E := [][]int{
		{1, 3},
		{2, 3},
		{3, 4},
		{3, 5},
	}
	expect := 4
	runSample(t, n, k, C, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 3
	C := []int{2, 1, 0, 1, 3, 2, 3, 2, 0, 3}
	E := [][]int{
		{1, 2},
		{3, 1},
		{3, 4},
		{5, 4},
		{4, 6},
		{3, 7},
		{8, 7},
		{9, 10},
		{9, 1},
	}
	expect := 7
	runSample(t, n, k, C, E, expect)
}
