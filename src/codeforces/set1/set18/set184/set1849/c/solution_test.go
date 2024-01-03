package main

import "testing"

func runSample(t *testing.T, s string, queries [][]int, expect int) {
	res := solve(s, queries)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101100"
	queries := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{5, 5},
		{1, 6},
	}
	expect := 3
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "100111"
	queries := [][]int{
		{2, 2},
		{1, 4},
		{1, 3},
		{1, 2},
		{1, 1},
	}
	expect := 3
	runSample(t, s, queries, expect)
}
