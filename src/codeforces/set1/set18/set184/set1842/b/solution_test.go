package main

import "testing"

func runSample(t *testing.T, stacks [][]int, x int, expect bool) {
	res := solve(stacks, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	stacks := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 5, 7, 9},
	}
	x := 7
	expect := true
	runSample(t, stacks, x, expect)
}
