package main

import "testing"

func runSample(t *testing.T, x int64, y int64, expect bool) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 4, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 8, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 34, 69, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 8935891487501725, 71487131900013807, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 10, 576460752303423487, false)
}

func TestSample6(t *testing.T) {
	runSample(t, 22, 11, true)
}
