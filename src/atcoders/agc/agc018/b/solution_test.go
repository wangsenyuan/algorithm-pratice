package main

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := solve(A)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{5, 1, 3, 4, 2},
		{2, 5, 3, 1, 4},
		{2, 3, 1, 4, 5},
		{2, 5, 4, 3, 1},
	}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{2, 1, 3},
		{2, 1, 3},
		{2, 1, 3},
	}
	expect := 3
	runSample(t, A, expect)
}
