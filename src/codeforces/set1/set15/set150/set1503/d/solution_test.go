package main

import "testing"

func runSample(t *testing.T, n int, P [][]int, expect int) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	P := [][]int{
		{3, 10},
		{6, 4},
		{1, 9},
		{5, 8},
		{2, 7},
	}
	expect := 2
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	P := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := -1
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	P := [][]int{
		{9, 8},
		{10, 7},
		{4, 6},
		{12, 11},
		{5, 3},
		{2, 1},
	}
	expect := -1
	runSample(t, n, P, expect)
}
