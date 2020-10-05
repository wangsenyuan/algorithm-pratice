package main

import "testing"

func runSample(t *testing.T, N int, V uint64, P []uint64, A0, B0, C0 uint64, M0 uint64, Q []uint64, A1, B1, C1 uint64, M1 uint64, expect int) {
	res := solve(N, V, P, A0, B0, C0, M0, Q, A1, B1, C1, M1)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, []uint64{1, 1}, 1, 1, 1, 1, []uint64{1, 1}, 1, 1, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, []uint64{1, 1}, 1, 1, 1, 2, []uint64{2, 1}, 1, 1, 1, 1, 729)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3, []uint64{1, 1}, 1, 1, 1, 2, []uint64{1, 1}, 0, 0, 0, 2, 387420489)
}
