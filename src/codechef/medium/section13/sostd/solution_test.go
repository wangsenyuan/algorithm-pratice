package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect int) {
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], B[i]}
	}
	res := solve(n, ps)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	B := []int{6, 5, 4}
	expect := 91
	runSample(t, n, A, B, expect)
}
