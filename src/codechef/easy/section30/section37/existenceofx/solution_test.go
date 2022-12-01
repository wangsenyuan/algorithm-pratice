package main

import "testing"

func runSample(t *testing.T, A, B, C int, expect bool) {
	res := solve(A, B, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 5, 7, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 13, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 7, 6, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 6, 6, true)
}
