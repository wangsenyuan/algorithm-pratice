package main

import "testing"

func runSample(t *testing.T, k int, A [][]int, D [][]byte, expect int) {
	res := solve(len(A), len(A[0]), k, A, D)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := [][]int{
		{4, 4, 4},
		{5, 4, 4},
		{4, 8, 4},
	}
	D := [][]byte{
		[]byte("DRR"),
		[]byte("DUU"),
		[]byte("URL"),
	}
	expect := 3
	runSample(t, k, A, D, expect)
}
