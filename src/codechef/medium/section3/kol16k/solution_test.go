package main

import "testing"

func runSample(t *testing.T, n int, P [][]int, expect int) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	P := [][]int{
		{1, 1},
		{4, 2},
		{9, 3},
	}
	expect := 3
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	P := [][]int{
		{1, 1},
		{4, 2},
		{5, 2},
		{9, 3},
	}
	expect := 3
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	P := [][]int{
		{2, 3},
		{4, 1},
		{1, 2},
		{3, 2},
	}
	expect := 4
	runSample(t, n, P, expect)
}
