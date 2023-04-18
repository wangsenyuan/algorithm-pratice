package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, x int, expect int) {
	res := solve(n, E, A, x)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 7, 1
	A := []int{1, 0, 1, 0, 1, 0, 1}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := 5
	runSample(t, n, E, A, x, expect)
}
