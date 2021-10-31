package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, q int, expect int) {
	res := solve(n, E, q)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 5, 1
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 3
	runSample(t, n, E, q, expect)
}

func TestSample2(t *testing.T) {
	n, q := 5, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	expect := 1
	runSample(t, n, E, q, expect)
}
