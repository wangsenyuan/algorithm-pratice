package main

import "testing"

func runSample(t *testing.T, A []int, B []int, S [][]int, expect bool) {
	res := solve(A, B, S)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 4, 2, 3}
	B := []int{3, 2, 5, 4, 1}
	S := [][]int{
		{1, 3},
		{2, 5},
	}
	expect := true
	runSample(t, A, B, S, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 5, 4, 2, 3}
	B := []int{3, 2, 5, 4, 1}
	S := [][]int{
		{1, 3},
		{2, 4},
	}
	expect := false
	runSample(t, A, B, S, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 5, 4, 2, 3}
	B := []int{3, 2, 5, 4, 1}
	S := [][]int{
		{1, 2},
		{2, 4},
	}
	expect := false
	runSample(t, A, B, S, expect)
}
