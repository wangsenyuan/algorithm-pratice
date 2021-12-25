package main

import "testing"

func runSample(t *testing.T, C, N, M int, S [][]int, expect int) {
	res := solve(C, N, M, S)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C, N, M := 2, 6, 3
	S := [][]int{
		{1, 2, 5, 6},
		{4, 6},
	}
	expect := 9
	runSample(t, C, N, M, S, expect)
}
