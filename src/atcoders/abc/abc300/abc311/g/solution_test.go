package main

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := solve(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{5, 4, 3},
		{4, 3, 2},
		{3, 2, 1},
	}
	expect := 48
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{3, 1, 4, 1, 5},
		{9, 2, 6, 5, 3},
		{5, 8, 9, 7, 9},
		{3, 2, 3, 8, 4},
	}
	expect := 231
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{1, 300, 300, 300, 300, 300},
		{300, 1, 300, 300, 300, 300},
		{300, 300, 1, 300, 300, 300},
		{300, 300, 300, 1, 300, 300},
		{300, 300, 300, 300, 1, 300},
		{300, 300, 300, 300, 300, 1},
	}
	expect := 810000
	runSample(t, mat, expect)
}
