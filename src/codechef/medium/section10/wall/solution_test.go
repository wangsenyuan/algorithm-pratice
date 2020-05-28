package main

import (
	"testing"
)

func runSample(t *testing.T, H, N, M, A, B, IND int, D []int, expect float64) {
	res := solve(H, N, M, A, B, IND, D)

	if res != int64(expect*2) {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H, N, M, A, B, IND := 5, 2, 2, 0, 1, 0
	D := []int{10, 10}
	expect := 25.0
	runSample(t, H, N, M, A, B, IND, D, expect)
}

func TestSample2(t *testing.T) {
	H, N, M, A, B, IND := 10, 10, 5, 2, 3, 4
	D := []int{2, 5, 1, 1, 4}
	expect := 140.0
	runSample(t, H, N, M, A, B, IND, D, expect)
}
