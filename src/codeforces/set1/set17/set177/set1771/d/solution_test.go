package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, s string, expect int) {
	res := solve(n, E, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abaca"
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
	}
	expect := 3
	runSample(t, n, E, s, expect)
}

func TestSample2(t *testing.T) {
	s := "caabadedb"
	n := 9
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{5, 6},
		{5, 7},
		{5, 8},
		{8, 9},
	}
	expect := 5
	runSample(t, n, E, s, expect)
}
