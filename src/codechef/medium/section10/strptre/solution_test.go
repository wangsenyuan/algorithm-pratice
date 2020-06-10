package main

import "testing"

func runSample(t *testing.T, D int, expect int) {
	res := solve(D)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 142857144)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 333333336)
}
