package main

import "testing"

func runSample(t *testing.T, n int, K int, A []int, expect bool) {
	res := solve(n, K, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 6, 0
	A := []int{1, 1, 1, 1, 1, 1}
	runSample(t, n, K, A, false)
}

func TestSample2(t *testing.T) {
	n, K := 5, 1
	A := []int{2, 4, 6, 3, 4}
	runSample(t, n, K, A, true)
}
