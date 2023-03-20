package main

import "testing"

func runSample(t *testing.T, k int, blocks [][]int, expect int) {
	res := solve(k, blocks)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 4
	blocks := [][]int{
		{1, 2, 1, 3},
		{2, 2, 2, 3},
		{3, 2, 3, 3},
		{4, 2, 4, 3},
		{1, 4, 2, 4},
		{2, 1, 3, 1},
		{2, 2, 3, 2},
		{4, 1, 4, 2},
	}
	expect := 7
	runSample(t, k, blocks, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	blocks := [][]int{
		{1, 1, 1, 2},
		{2, 1, 2, 2},
		{1, 1, 1, 2},
		{1, 1, 2, 1},
		{1, 2, 2, 2},
		{1, 1, 2, 1},
		{1, 2, 2, 2},
	}
	expect := 4
	runSample(t, k, blocks, expect)
}

func TestSample3(t *testing.T) {
	k := 6
	blocks := [][]int{
		{3, 3, 3, 4},
	}
	expect := 1
	runSample(t, k, blocks, expect)
}
