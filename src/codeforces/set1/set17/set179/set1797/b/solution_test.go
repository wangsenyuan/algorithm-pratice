package main

import "testing"

func runSample(t *testing.T, mat [][]int, k int, expect bool) {
	res := solve(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 0, 1, 1},
		{1, 0, 0, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
	}
	k := 3
	expect := true
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{0},
	}
	k := 1
	expect := true
	runSample(t, mat, k, expect)
}
