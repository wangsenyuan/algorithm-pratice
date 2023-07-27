package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int, expect int) {
	res := solve(n, E, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{2, 2, 2}
	E := [][]int{
		{1, 2},
		{3, 2},
	}
	expect := 3
	runSample(t, n, E, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 5
	runSample(t, n, E, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{0, 0, 0, -1}
	E := [][]int{
		{3, 1},
		{4, 1},
		{2, 4},
	}
	expect := 2
	runSample(t, n, E, A, expect)
}
