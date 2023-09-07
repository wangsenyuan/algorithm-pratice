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
		{2, 3},
		{4, 5},
	}
	expect := 2
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{4, 4, 3},
		{2, 1, 4},
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample4(t *testing.T) {
	mat := [][]int{
		{1, 9, 4, 6, 5, 8},
		{10, 9, 5, 8, 11, 6},
		{24, 42, 32, 8, 11, 1},
		{23, 1, 9, 69, 13, 3},
		{13, 22, 60, 12, 14, 17},
	}
	expect := 3
	runSample(t, mat, expect)
}
