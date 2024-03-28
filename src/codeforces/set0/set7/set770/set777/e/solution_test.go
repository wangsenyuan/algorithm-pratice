package main

import "testing"

func runSample(t *testing.T, rings [][]int, expect int) {
	res := solve(rings)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rings := [][]int{
		{1, 5, 1},
		{2, 6, 2},
		{3, 7, 3},
	}
	expect := 6
	runSample(t, rings, expect)
}

func TestSample2(t *testing.T) {
	rings := [][]int{
		{1, 2, 1},
		{1, 3, 3},
		{4, 6, 2},
		{5, 7, 1},
	}
	expect := 4
	runSample(t, rings, expect)
}
