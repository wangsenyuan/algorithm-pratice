package main

import "testing"

func runSample(t *testing.T, N, K, x int, A []int, expect int64) {
	res := solve(N, K, x, A)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, K, x, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, x := 4, 3, 4
	A := []int{2, 1, 5}
	runSample(t, N, K, x, A, 12)
}

func TestSample2(t *testing.T) {
	N, K, x := 2, 2, 9
	A := []int{3, 6}
	runSample(t, N, K, x, A, 9)
}

func TestSample3(t *testing.T) {
	N, K, x := 10, 3, 49
	A := []int{1, 43, 44}
	runSample(t, N, K, x, A, 415)
}
