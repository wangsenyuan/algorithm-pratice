package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, B []int, expect int) {
	res := solve(n, E, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 1, 0, 0}
	B := []int{1, 1, 1, 0}
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 1
	runSample(t, n, E, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 1, 1, 0, 0}
	B := []int{1, 0, 1, 1, 1}
	E := [][]int{
		{5, 3},
		{3, 1},
		{2, 1},
		{4, 3},
	}
	expect := 2
	runSample(t, n, E, A, B, expect)
}
