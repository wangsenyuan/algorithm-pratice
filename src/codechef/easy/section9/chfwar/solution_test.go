package main

import "testing"

func runSample(t *testing.T, N int, A []int, F int, a, b int) {
	c, d := solve(N, A, F)

	if a != c || d != b {
		t.Errorf("Sample %d %v %d, expect %d, %d, but got %d, %d", N, A, F, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	A := []int{12, 34, 45, 5}
	F := 10
	a, b := 4, 100
	runSample(t, N, A, F, a, b)
}

func TestSample2(t *testing.T) {
	N := 5
	A := []int{10, 15, 43, 20}
	F := 5
	a, b := -1, -1
	runSample(t, N, A, F, a, b)
}
