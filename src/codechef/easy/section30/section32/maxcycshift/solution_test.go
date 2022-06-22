package main

import "testing"

func runSample(t *testing.T, A []uint64, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []uint64{0, 0}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []uint64{1, 1, 1, 2, 2, 2}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []uint64{1, 2, 2, 8}
	expect := 4
	runSample(t, A, expect)
}
