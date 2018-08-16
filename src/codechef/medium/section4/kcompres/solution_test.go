package main

import "testing"

func runSample(t *testing.T, N, S int, A []int, expect int) {
	res := solve(N, S, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, S, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, S := 4, 8
	A := []int{2, 7, 1, 12}
	expect := 2
	runSample(t, N, S, A, expect)
}

func TestSample2(t *testing.T) {
	N, S := 8, 20
	A := []int{4, 2, 8, 1, 4, 3, 8, 1}
	expect := 4
	runSample(t, N, S, A, expect)
}
