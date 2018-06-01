package main

import "testing"

func runSample(t *testing.T, n uint64, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 6, 4)
}

func TestSample6(t *testing.T) {
	runSample(t, 7, 5)
}
