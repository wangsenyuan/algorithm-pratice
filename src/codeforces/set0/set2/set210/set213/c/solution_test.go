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
		{5},
	}
	expect := 5
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{11, 14},
		{16, 12},
	}
	expect := 53
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{25, 16, 25},
		{12, 18, 19},
		{11, 13, 8},
	}
	expect := 136
	runSample(t, mat, expect)
}
