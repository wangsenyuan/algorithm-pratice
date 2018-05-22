package main

import "testing"

func runSample(t *testing.T, N, M int, MTR [][]byte, expect int64) {
	res := solve(N, M, MTR)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, MTR, expect, res)
	}
}

func TestSample(t *testing.T) {
	N, M := 2, 3
	MTR := [][]byte{
		[]byte("011"),
		[]byte("111"),
	}
	runSample(t, N, M, MTR, 6)
}

func TestSample1(t *testing.T) {
	N, M := 2, 3
	MTR := [][]byte{
		[]byte("011"),
		[]byte("110"),
	}
	runSample(t, N, M, MTR, 4)
}
