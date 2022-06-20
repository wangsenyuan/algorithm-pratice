package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12345, 51234)
}

func TestSample2(t *testing.T) {
	runSample(t, 54321, 54321)
}

func TestSample3(t *testing.T) {
	runSample(t, 10901, 11090)
}

func TestSample4(t *testing.T) {
	runSample(t, 211011, 211011)
}

func TestSample5(t *testing.T) {
	runSample(t, 90, 9)
}