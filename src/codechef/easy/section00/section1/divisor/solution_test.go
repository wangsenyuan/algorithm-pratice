package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expec %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 64, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 30, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 120, 1)
}
