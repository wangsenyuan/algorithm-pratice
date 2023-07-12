package main

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := int(solve(mat))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{4, 2},
		{2, 4},
		{4, 2},
		{2, 4},
	}
	expect := 8
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 18},
	}
	expect := 42
	runSample(t, mat, expect)
}
