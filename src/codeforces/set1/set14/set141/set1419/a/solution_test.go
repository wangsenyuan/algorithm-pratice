package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, "2", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, "3", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, "102", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, "2049", 2)
}
