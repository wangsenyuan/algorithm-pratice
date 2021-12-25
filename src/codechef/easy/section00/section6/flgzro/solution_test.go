package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, expect int) {
	res := solve(n, A, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	A := []int{1}
	expect := 1
	runSample(t, n, nil, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	A := []int{1, 1, 1, 2}
	expect := 4
	runSample(t, n, E, A, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	A := []int{1, 2, 3, 4, 1, 2, 1}
	expect := 14
	runSample(t, n, E, A, expect)
}
