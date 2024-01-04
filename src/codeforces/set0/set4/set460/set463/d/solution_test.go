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
		{1, 4, 2, 3},
		{4, 1, 2, 3},
		{1, 2, 4, 3},
	}
	expect := 3
	runSample(t, mat, expect)
}
