package main

import "testing"

func runSample(t *testing.T, x, y, k int, expect int64) {
	res := solve(x, y, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 5, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, 42, 13, 24, 33)
}

func TestSample3(t *testing.T) {
	runSample(t, 12, 11, 12, 25)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000, 1000000000, 1000000000, 2000000003)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 1000000000, 1000000000, 1000000001999999999)
}
