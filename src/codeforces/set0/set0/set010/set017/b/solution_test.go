package main

import "testing"

func runSample(t *testing.T, P []int, E [][]int, expect int) {
	res := solve(len(P), P, E)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{7, 2, 3, 1}
	E := [][]int{
		{1, 2, 5},
		{2, 4, 1},
		{3, 4, 1},
		{1, 3, 5},
	}
	expect := 11
	runSample(t, P, E, expect)
}

func TestSample2(t *testing.T) {
	P := []int{1, 2, 3}
	E := [][]int{
		{3, 1, 2},
		{3, 1, 3},
	}
	expect := -1
	runSample(t, P, E, expect)
}

func TestSample3(t *testing.T) {
	P := []int{3, 9, 2, 1, 8}
	E := [][]int{
		{2, 5, 10},
		{1, 3, 8},
		{3, 4, 9},
		{5, 4, 2},
		{2, 1, 4},
		{5, 1, 4},
		{2, 4, 2},
		{1, 4, 7},
		{5, 1, 2},
	}
	expect := 22
	runSample(t, P, E, expect)
}
