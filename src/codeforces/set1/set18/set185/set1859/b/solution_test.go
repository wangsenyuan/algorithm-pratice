package main

import "testing"

func runSample(t *testing.T, a [][]int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{1001, 7, 1007, 5},
		{8, 11, 6},
		{2, 9},
	}
	expect := 19
	runSample(t, a, expect)
}
