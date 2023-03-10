package main

import "testing"

func runSample(t *testing.T, n int, k int, p int, expect int) {
	res := solve(n, k, p)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 1, 1000000, 55)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 2, 1000000, 385)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 3, 1000000, 3025)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 4, 1000000, 25333)
}
