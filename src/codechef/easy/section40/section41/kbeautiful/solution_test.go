package main

import "testing"

func runSample(t *testing.T, K int, M int, A []int64, expect int) {
	res := solve(K, M, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	M, K := 4, 3
	A := []int64{1, 2, 1, 2, 1}
	expect := 5
	runSample(t, K, M, A, expect)
}
