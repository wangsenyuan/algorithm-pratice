package main

import "testing"

func runSample(t *testing.T, x, s int, expect int) {
	res := solve(x, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 23, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 13, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 25, 25)
}
