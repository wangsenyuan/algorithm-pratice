package main

import "testing"

func runSample(t *testing.T, n, m int, P []int, A [][]int, expect int) {
	res := solve(n, m, P, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 12, 8
	P := []int{4, 1, 1, 2, 7, 8, 4, 2, 12, 12, 4}
	A := [][]int{
		{7},
		{7},
		{6, 11},
		{8, 11},
		{10, 8},
		{12, 8},
		{8, 10},
		{7, 12},
	}
	expect := 6
	runSample(t, n, m, P, A, expect)
}
