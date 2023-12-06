package main

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{4, 9, 4, 6, 8},
		{7},
		{8, 6},
		{1},
	}
	expect := 4
	runSample(t, A, expect)
}
