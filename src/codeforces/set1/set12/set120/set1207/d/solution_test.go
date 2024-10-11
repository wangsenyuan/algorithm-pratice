package main

import "testing"

func runSample(t *testing.T, pairs [][]int, expect int) {
	res := solve(pairs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{1, 1},
		{2, 2},
		{3, 1},
	}
	expect := 3
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{2, 3},
		{2, 2},
		{2, 1},
		{2, 4},
	}
	expect := 0
	runSample(t, pairs, expect)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{
		{1, 1},
		{1, 1},
		{2, 3},
	}
	expect := 4
	runSample(t, pairs, expect)
}
