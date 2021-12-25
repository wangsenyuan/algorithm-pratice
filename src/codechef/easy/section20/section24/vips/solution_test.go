package main

import "testing"

func runSample(t *testing.T, A []int, E [][]int, expect int) {
	res := int(solve(len(A), A, E))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {

	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	A := []int{2, 2, 2, 2}
	expect := 6
	runSample(t, A, E, expect)
}

func TestSample2(t *testing.T) {

	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}
	A := []int{5, 2, 3, 1, 4}
	expect := 7
	runSample(t, A, E, expect)
}
