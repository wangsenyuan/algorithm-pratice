package main

import "testing"

func runSample(t *testing.T, A, B int, L, R int64, expect int64) {
	res := solve(A, B, L, R)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 15, 1, 10, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 235714494, 193478285, 2, 100000000000, 65435081000)
}
