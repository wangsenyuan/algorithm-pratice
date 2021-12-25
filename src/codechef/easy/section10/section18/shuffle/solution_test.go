package main

import "testing"

func runSample(t *testing.T, N int, A []int, K int, expect bool) {
	res := solve(N, A, K)

	if res != expect {
		t.Errorf("Sample %d %v %d, expect %t, but got %t", N, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 4
	A := []int{1, 4, 2, 3}
	K := 2
	runSample(t, N, A, K, false)
}

func TestSample2(t *testing.T) {
	N := 5
	A := []int{1, 2, 3, 4, 5}
	K := 3
	runSample(t, N, A, K, true)
}

func TestSample3(t *testing.T) {
	N := 5
	A := []int{1, 3, 2, 4, 5}
	K := 3
	runSample(t, N, A, K, false)
}
