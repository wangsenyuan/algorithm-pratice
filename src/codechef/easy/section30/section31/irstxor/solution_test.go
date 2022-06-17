package main

import "testing"

func runSample(t *testing.T, c int, expect int64) {
	res := solve(c)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 70)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 105)
}
