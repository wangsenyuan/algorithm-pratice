package main

import "testing"

func runSample(t *testing.T, n int, mat [][]int, expect bool) {
	res := solve(n, mat)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	mat := [][]int{
		{1, 2, 3, 4},
		{2, 3, 1, 4},
		{3, 2, 1, 4},
		{4, 2, 3, 1},
	}
	expect := true
	runSample(t, n, mat, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	mat := [][]int{
		{1, 3, 5, 2, 4, 6},
		{6, 3, 5, 2, 1, 4},
	}
	expect := true
	runSample(t, n, mat, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	mat := [][]int{
		{3, 1, 2},
		{2, 3, 1},
		{1, 3, 2},
	}
	expect := true
	runSample(t, n, mat, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	mat := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
	expect := false
	runSample(t, n, mat, expect)
}
