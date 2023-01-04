package main

import "testing"

func runSample(t *testing.T, lines [][]int, expect bool) {
	res := solve(lines)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lines := [][]int{
		{0, 0, 0, 2},
		{2, 0, 0, 0},
		{2, 2, 2, 0},
		{0, 2, 2, 2},
	}
	expect := true
	runSample(t, lines, expect)
}
