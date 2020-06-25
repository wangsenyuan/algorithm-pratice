package main

import "testing"

func runSample(t *testing.T, N int, M int, A []int, expect int) {
	res := solve(N, M, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 3, 3
	A := []int{1, 2, 4}
	expect := 3
	runSample(t, N, M, A, expect)
}
