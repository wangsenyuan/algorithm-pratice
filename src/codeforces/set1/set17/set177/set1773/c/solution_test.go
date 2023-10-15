package main

import "testing"

func runSample(t *testing.T, blocks [][]int, expect []int) {
	split, merge := solve(blocks)

	if expect[0] != split || expect[1] != merge {
		t.Fatalf("Sample expect %v, but got %d %d", expect, split, merge)
	}
}

func TestSample1(t *testing.T) {
	blocks := [][]int{
		{3, 5, 8},
		{9, 2},
	}
	expect := []int{1, 2}
	runSample(t, blocks, expect)
}
