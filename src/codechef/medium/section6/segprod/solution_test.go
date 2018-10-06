package main

import (
	"testing"
)

func runSample(t *testing.T, N, P, Q int, A []int, B []int, expect int) {
	res := solve(N, P, Q, A, B)
	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", N, P, Q, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, P, Q := 6, 113, 3
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{1, 4}
	runSample(t, N, P, Q, A, B, 8)
}

func TestSample2(t *testing.T) {
	N, P, Q := 6, 113, 70
	A := []int{1, 2, 3, 4, 5, 6}
	B := []int{1, 4, 3}
	runSample(t, N, P, Q, A, B, 21)
}
