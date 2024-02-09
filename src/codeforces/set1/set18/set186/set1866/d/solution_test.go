package main

import "testing"

func runSample(t *testing.T, mat [][]int, k int, expect int) {
	res := solve(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	mat := [][]int{
		{10, 4, 2},
		{8, 1, 9},
		{4, 8, 2},
	}
	expect := 27
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	mat := [][]int{
		{5, 9, 4},
		{1, 3, 1},
		{2, 8, 7},
	}
	expect := 17
	runSample(t, mat, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	mat := [][]int{
		{5, 9, 10, 1},
		{1, 3, 1, 5},
		{2, 5, 7, 2},
	}
	expect := 19
	runSample(t, mat, k, expect)
}
