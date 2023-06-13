package main

import "testing"

func runSample(t *testing.T, n int, k int, parts [][]int, expect int) {
	res := solve(n, k, parts)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 10, 3
	parts := [][]int{
		{1, 3},
		{2, 4},
		{6, 9},
		{6, 9},
		{1, 8},
	}
	expect := 14
	runSample(t, n, k, parts, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 3
	parts := [][]int{
		{2, 4},
		{4, 6},
		{3, 5},
	}
	expect := 8
	runSample(t, n, k, parts, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 1
	parts := [][]int{
		{3, 3},
		{1, 1},
		{2, 2},
		{4, 4},
	}
	expect := 2
	runSample(t, n, k, parts, expect)
}
