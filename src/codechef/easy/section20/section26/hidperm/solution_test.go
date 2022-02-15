package main

import "testing"

func runSample(t *testing.T, P [][]int, expect int) {
	res := solve(len(P), P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := [][]int{
		{2, 2},
		{1, 2},
	}
	expect := 1
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := [][]int{
		{3, 3},
		{1, 3},
		{2, 2},
	}
	expect := 0
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := [][]int{
		{1, 1},
		{5, 5},
		{1, 5},
		{1, 2},
		{4, 5},
	}
	expect := 6
	runSample(t, P, expect)
}

func TestSample4(t *testing.T) {
	P := [][]int{
		{1, 1},
		{7, 7},
		{3, 3},
		{1, 7},
		{6, 7},
		{1, 4},
		{3, 4},
	}
	expect := 45
	runSample(t, P, expect)
}
