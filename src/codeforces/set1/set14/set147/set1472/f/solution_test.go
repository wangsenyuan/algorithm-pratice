package main

import "testing"

func runSample(t *testing.T, n int, blocks [][]int, expect bool) {
	res := solve(n, blocks)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	blocks := [][]int{
		{2, 2},
		{1, 4},
	}
	expect := true
	runSample(t, n, blocks, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	blocks := [][]int{
		{2, 1},
		{2, 3},
	}
	expect := false
	runSample(t, n, blocks, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	blocks := [][]int{
		{2, 1},
		{2, 3},
		{2, 4},
		{2, 6},
	}
	expect := false
	runSample(t, n, blocks, expect)
}
