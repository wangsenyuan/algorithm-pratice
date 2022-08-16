package main

import "testing"

func runSample(t *testing.T, A, B [][]int, expect bool) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	B := [][]int{
		{1, 4, 3},
		{6, 5, 2},
	}
	expect := true
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{12, 11, 8},
		{7, 1, 1},
		{9, 2, 4},
	}
	B := [][]int{
		{4, 1, 8},
		{2, 1, 11},
		{9, 7, 12},
	}
	expect := true
	runSample(t, A, B, expect)
}
