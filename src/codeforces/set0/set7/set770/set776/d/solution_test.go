package main

import "testing"

func runSample(t *testing.T, n int, locked []int, m int, lines [][]int, expect bool) {
	res := solve(n, locked, m, lines)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	locked := []int{1, 0, 1}
	lines := [][]int{
		{1, 3},
		{1, 2},
		{2, 3},
	}
	expect := false
	runSample(t, n, locked, m, lines, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	locked := []int{1, 0, 1}
	lines := [][]int{
		{1, 2, 3},
		{2},
		{1, 3},
	}
	expect := true
	runSample(t, n, locked, m, lines, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	locked := []int{1, 0, 1}
	lines := [][]int{
		{1, 2, 3},
		{1, 2},
		{3},
	}
	expect := false
	runSample(t, n, locked, m, lines, expect)
}
