package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "100010", "0", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "0", "100010", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "11100", "1010", 3)
}
