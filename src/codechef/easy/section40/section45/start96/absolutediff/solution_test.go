package main

import "testing"

func runSample(t *testing.T, n int, A [][]int, B [][]int, expect bool) {
	res := solve(n, A, B)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := [][]int{
		{1, -1},
		{3, 3},
	}
	B := [][]int{
		{1, 2, 2},
		{2, 3, 2},
	}
	expect := true
	runSample(t, n, A, B, expect)
}
