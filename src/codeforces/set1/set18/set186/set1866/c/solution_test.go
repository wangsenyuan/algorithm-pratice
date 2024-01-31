package main

import "testing"

func runSample(t *testing.T, n int, edges [][][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][][]int{
		{
			{4, 0},
			{3, 1},
		},
		{},
		{
			{2, 0},
		},
		{
			{3, 1},
			{5, 1},
		},
		{},
	}
	expect := 4
	runSample(t, n, edges, expect)
}
