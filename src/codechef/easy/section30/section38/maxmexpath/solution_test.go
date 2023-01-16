package main

import "testing"

func runSample(t *testing.T, n int, A []int, E [][]int, expect int) {
	res := solve(n, A, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{0, 1, 2}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 2
	runSample(t, n, A, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{0, 1, 2, 3}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
	}
	expect := 3
	runSample(t, n, A, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 0}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
	}
	expect := 2
	runSample(t, n, A, E, expect)
}
