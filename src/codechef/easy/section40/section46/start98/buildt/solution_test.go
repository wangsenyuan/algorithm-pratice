package main

import "testing"

func runSample(t *testing.T, A []int, E [][]int, expect int) {
	res := solve(A, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 0, 1, 0}
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := 1
	runSample(t, A, E, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 0}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 3
	runSample(t, A, E, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 0, 1}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := 3
	runSample(t, A, E, expect)
}
