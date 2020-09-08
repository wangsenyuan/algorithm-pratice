package main

import "testing"

func runSample(t *testing.T, N, K, D, m int, A []int, expect int64) {
	res := solve(N, K, D, m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, D, m := 5, 2, 2, 10
	A := []int{1, 2, 3, 4, 5}
	var expect int64 = 110
	runSample(t, N, K, D, m, A, expect)
}
