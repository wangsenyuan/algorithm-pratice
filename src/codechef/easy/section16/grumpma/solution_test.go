package main

import "testing"

func runSample(t *testing.T, n, K, M int, A []int, expect int) {
	res := solve(n, K, M, A)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", n, K, M, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
		n, K, M := 12, 4, 3
		A := []int{4, 5, 6, 7, 1, 4, 6, 9, 0, 0, 10, 2}
		expect := 8
		runSample(t, n, K, M, A, expect)
}