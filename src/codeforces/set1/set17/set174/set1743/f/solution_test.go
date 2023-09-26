package main

import "testing"

func runSample(t *testing.T, ss [][]int, expect int) {
	res := solve(ss)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ss := [][]int{
		{3, 5},
		{4, 8},
		{2, 2},
		{1, 9},
	}
	expect := 162
	runSample(t, ss, expect)
}
