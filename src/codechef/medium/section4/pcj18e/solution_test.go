package main

import "testing"

func runSample(t *testing.T, N int, A []int, expect int) {
	res := solve(N, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	A := []int{1, 2, 3, 4, 5}
	runSample(t, N, A, 0)
}

func TestSample2(t *testing.T) {
	N := 5
	A := []int{1, 3, 2, 5, 4}
	runSample(t, N, A, 3)
}
