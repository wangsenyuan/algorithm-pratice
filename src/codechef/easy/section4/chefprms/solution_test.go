package main

import "testing"

func runSample(t *testing.T, N int, expect bool) {
	res := solve(N)
	if res != expect {
		t.Errorf("Sample %d, expect %t, but got %t", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 30, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 45, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 62, false)
}
