package main

import "testing"

func runSample(t *testing.T, n int, val []int, Q int, in []int64, M int64, A int, B int) {
	m, a, b := solve(n, val, Q, in)

	if m != M || a != A || b != B {
		t.Errorf("Sample expect %d %d %d, but got %d %d %d", M, A, B, m, a, b)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	val := []int{3, 5}
	Q := 2
	in := []int64{8, 1, 10}
	var M int64 = 7
	A, B := 1, 1
	runSample(t, n, val, Q, in, M, A, B)
}
