package main

import "testing"

func runSample(t *testing.T, n int, events [][]int, expect int) {
	res := solve(n, events)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, events, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 4
	events := [][]int{
		{2, 5},
		{9, 7},
		{15, 6},
		{9, 3},
	}

	expect := 3
	runSample(t, n, events, expect)
}
