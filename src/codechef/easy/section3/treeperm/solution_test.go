package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, S int, E [][]int, expect int) {
	res := solve(n, A, B, S, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, S := 3, 2
	E := [][]int{
		{1, 2},
		{2, 3},
	}

	A := []int{4, 5, 6}
	B := []int{4, 6, 5}
	expect := 2
	runSample(t, n, A, B, S, E, expect)
}

func TestSample2(t *testing.T) {
	n, S := 6, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 5},
		{3, 6},
	}

	A := []int{10, 20, 30, 40, 50, 60}
	B := []int{10, 40, 50, 20, 30, 60}
	expect := 3
	runSample(t, n, A, B, S, E, expect)
}
