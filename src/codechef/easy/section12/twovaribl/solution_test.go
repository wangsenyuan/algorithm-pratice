package main

import "testing"

func runSample(t *testing.T, x int, expect int) {
	res := solve(x)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but %d", x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 10000000, 46)
}

func TestSample5(t *testing.T) {
	runSample(t, 100000000, 52)
}
