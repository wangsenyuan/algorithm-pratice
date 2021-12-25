package main

import "testing"

func runSample(t *testing.T, n int, A []uint64, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []uint64{1, 2, 7}
	expect := false
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []uint64{1, 2}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []uint64{6, 5, 8}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []uint64{12, 32, 45, 23, 47}
	expect := false
	runSample(t, n, A, expect)
}
