package main

import "testing"

func runSample(t *testing.T, N, K, L int, A []int, expect int) {
	res := solve(N, K, L, A)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, K, L, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, L := 2, 2, 1
	A := []int{1, 2}
	expect := 3
	runSample(t, N, K, L, A, expect)
}

func TestSample2(t *testing.T) {
	N, K, L := 3, 3, 2
	A := []int{1, 2, 2}
	expect := 11
	runSample(t, N, K, L, A, expect)
}
