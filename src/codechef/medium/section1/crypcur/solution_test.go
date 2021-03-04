package main

import "testing"

func runSample(t *testing.T, n, m int, C [][]int, expect int) {
	res := solve(n, m, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 6
	C := [][]int{
		{1, 2, 1},
		{2, 3, 3},
		{3, 1, 3},
		{1, 3, 1},
		{3, 2, 1},
		{3, 2, 5},
	}
	expect := 1
	runSample(t, n, m, C, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 3
	C := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{2, 4, 1},
	}
	expect := -1
	runSample(t, n, m, C, expect)
}
