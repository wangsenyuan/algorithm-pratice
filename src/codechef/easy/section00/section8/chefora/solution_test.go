package main

import "testing"

func runSample(t *testing.T, L, R int, expect int) {
	res := solve(L, R)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 11, 541416750)
}
