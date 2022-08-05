package main

import "testing"

func runSample(t *testing.T, A, B int64, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample %d %d expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 8, 36)
}

func TestSample2(t *testing.T) {
	runSample(t, 28, 138, 495)
}

func TestSample3(t *testing.T) {
	runSample(t, 314159, 314159, 7)
}

func TestSample5(t *testing.T) {
	runSample(t, 0, 20, 94)
}
