package main

import "testing"

func runSample(t *testing.T, X int, A, B [][]int, expect bool) {
	res := solve(X, A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := 2
	A := [][]int{
		{1, 2},
		{0, 1},
	}
	B := [][]int{
		{0, 0},
		{0, 0},
	}
	expect := true
	runSample(t, X, A, B, expect)
}

func TestSample2(t *testing.T) {
	X := 2
	A := [][]int{
		{1, 2},
		{0, 1},
	}
	B := [][]int{
		{0, 0},
		{-1, 0},
	}
	expect := false
	runSample(t, X, A, B, expect)
}

func TestSample3(t *testing.T) {
	X := 2
	A := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	B := [][]int{
		{1, 0},
		{2, 0},
		{3, 0},
	}
	expect := false
	runSample(t, X, A, B, expect)
}
