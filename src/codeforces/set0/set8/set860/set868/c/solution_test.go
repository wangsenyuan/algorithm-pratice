package main

import "testing"

func runSample(t *testing.T, k int, problems [][]int, expect bool) {
	res := solve(k, problems)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	problems := [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
	}
	expect := true
	runSample(t, k, problems, expect)
}
