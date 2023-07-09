package main

import "testing"

func runSample(t *testing.T, mat [][]int, k int, expect int) {
	res := solve(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 2, 3, 4},
		{5, 2, 2, 2},
		{7, 1, 1, 4},
	}
	k := 3
	expect := 24
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 2, 4, 2, 1},
		{3, 5, 1, 2, 4},
		{1, 5, 7, 1, 2},
		{3, 8, 7, 1, 2},
		{8, 4, 7, 1, 6},
	}
	k := 4
	expect := 56
	runSample(t, mat, k, expect)
}
