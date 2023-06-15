package main

import "testing"

func runSample(t *testing.T, A [][]int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	var expect int64 = 13
	runSample(t, A, expect)
}
