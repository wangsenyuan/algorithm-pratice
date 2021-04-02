package main

import "testing"

func runSample(t *testing.T, A, B, C, D []int, X, Y, Z [][]int, expect int) {
	res := solve(A, B, C, D, X, Y, Z)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := []int{5, 6, 7}
	C := []int{8, 9}
	D := []int{10}
	X := [][]int{
		{1, 2},
		{1, 1},
	}
	Y := [][]int{
		{3, 1},
		{3, 2},
	}
	Z := [][]int{
		{1, 1},
	}
	expect := 26
	runSample(t, A, B, C, D, X, Y, Z, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	B := []int{1}
	C := []int{1}
	D := []int{1}
	X := [][]int{
		{1, 1},
	}
	Y := [][]int{

	}
	Z := [][]int{
	}
	expect := -1
	runSample(t, A, B, C, D, X, Y, Z, expect)
}
