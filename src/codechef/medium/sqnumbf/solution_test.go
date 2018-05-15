package main

import "testing"

func runSample(t *testing.T, n int, A []int64) {
	x := solve(n, A)

	var product int64 = 1

	for i := 0; i < n; i++ {
		product *= A[i]
	}

	if x < 0 || product%(x*x) != 0 {
		t.Errorf("sample %d %v, answer %d is wrong", n, A, x)
	}
}

func TestSample(t *testing.T) {
	n := 3
	A := []int64{21, 11, 6}
	runSample(t, n, A)
}
