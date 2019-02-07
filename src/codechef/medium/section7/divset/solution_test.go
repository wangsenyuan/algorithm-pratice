package main

import "testing"

func runSample(t *testing.T, N, K, C int, A []int, expect int) {
	res := solve(N, K, C, A)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, K, C, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, C := 6, 3, 2
	A := []int{4, 1, 2, 2, 3, 1}
	expect := 1
	runSample(t, N, K, C, A, expect)
}

func TestSample2(t *testing.T) {
	N, K, C := 6, 3, 2
	A := []int{1, 2, 2, 1, 4, 4}
	expect := 2
	runSample(t, N, K, C, A, expect)
}
