package main

import "testing"

const X = 300000

func runSample(t *testing.T, n int) {
	A, B := solve(n)

	for i := 0; i < n; i++ {
		if A[i] > X || B[i] > X {
			t.Fatalf("sample fail as exceeds at %d, %d %d", i, A[i], B[i])
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 1 << 17)
}