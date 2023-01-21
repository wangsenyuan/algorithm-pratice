package main

import "testing"

func runSample(t *testing.T, n, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 1, "100010", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 2, "000000", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 1, "10101", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 1, "001", 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 2, "00", 1)
}

func TestSample6(t *testing.T) {
	runSample(t, 1, 1, "0", 1)
}
