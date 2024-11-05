package main

import "testing"

func runSample(t *testing.T, n int, lines [][]int, expect int) {
	res := solve(n, lines)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	lines := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	expect := 16
	runSample(t, n, lines, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	lines := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := 24
	runSample(t, n, lines, expect)
}
