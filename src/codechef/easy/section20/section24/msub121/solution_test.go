package main

import "testing"

func runSample(t *testing.T, A []int, Q [][]int, expect int) {
	res := solve(len(A), A, Q)

	if len(res) != expect {
		t.Errorf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 2, 8, 3, 7, 6}
	Q := [][]int{
		{3, 7},
		{3, 3},
		{3, 8},
		{8, 7},
		{8, 6},
	}
	expect := 3
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 1, 2, 1}
	Q := [][]int{
		{2, 1},
		{3, 1},
		{1, 1},
		{3, 2},
		{1, 2},
	}
	expect := 5
	runSample(t, A, Q, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1, 4}
	Q := [][]int{
		{1, 3},
		{1, 1},
		{1, 4},
		{2, 4},
		{4, 4},
	}
	expect := 3
	runSample(t, A, Q, expect)
}
