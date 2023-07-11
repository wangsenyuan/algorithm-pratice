package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect int) {
	res := solve(n, E, Q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2, 5},
		{2, 3, 7},
		{2, 4, 4},
		{4, 5, 2},
		{4, 6, 8},
	}
	Q := [][]int{
		{1, 6},
		{5, 3},
	}
	expect := 22
	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 5},
		{2, 3, 4},
		{1, 4, 3},
		{4, 3, 7},
		{3, 5, 2},
	}
	Q := [][]int{
		{1, 5},
		{1, 3},
		{3, 3},
		{1, 5},
	}
	expect := 13
	runSample(t, n, E, Q, expect)
}
