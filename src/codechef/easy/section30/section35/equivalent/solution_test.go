package main

import "testing"

func runSample(t *testing.T, A, B int, expect bool) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 4, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 12, 24, false)
}
