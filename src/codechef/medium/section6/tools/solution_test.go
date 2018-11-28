package main

import "testing"

func runSample(t *testing.T, n int, lines [][]int, expect int) {
	res := solve(n, lines)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, lines, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	lines := [][]int{
		{1, 0, 0, 1},
		{0, 0, 1, 1},
	}
	expect := 4
	runSample(t, n, lines, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	lines := [][]int{
		{0, 3, 0, 1},
		{0, 4, 0, 2},
		{0, 5, 0, 3},
	}
	expect := 10
	runSample(t, n, lines, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	lines := [][]int{
		{0, 1, 0, 2},
		{0, 1, 0, 2},
		{0, 1, 0, 2},
	}
	expect := 6
	runSample(t, n, lines, expect)
}
