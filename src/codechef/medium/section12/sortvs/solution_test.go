package main

import "testing"

func runSample(t *testing.T, n int, m int, P []int, X [][]int, expect int) {
	res := solve(n, m, P, X)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 1
	P := []int{2, 3, 1}
	X := [][]int{{2, 3}}
	expect := 1
	runSample(t, n, m, P, X, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 10
	P := []int{2, 4, 5, 1, 3}
	X := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 4},
		{3, 5},
		{4, 5},
	}

	expect := 0
	runSample(t, n, m, P, X, expect)
}

func TestSample3(t *testing.T) {
	n, m := 4, 1
	P := []int{3, 1, 4, 2}
	X := [][]int{
		{1, 2},
	}

	expect := 2
	runSample(t, n, m, P, X, expect)
}
