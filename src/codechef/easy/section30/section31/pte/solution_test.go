package main

import "testing"

func runSample(t *testing.T, n int, A []int, E [][]int, expect int64) {
	res := solve(n, A, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	A := []int{15, 30, 15, 5, 3, 15, 3, 3, 5, 5}
	E := [][]int{
		{1, 2},
		{1, 5},
		{2, 3},
		{2, 4},
		{5, 6},
		{5, 7},
		{5, 8},
		{7, 9},
		{7, 10},
	}

	var expect int64 = 33
	runSample(t, n, A, E, expect)
}
