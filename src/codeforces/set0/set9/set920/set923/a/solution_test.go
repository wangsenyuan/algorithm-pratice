package main

import "testing"

func runSample(t *testing.T, x2 int, expect int) {
	res := solve(x2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 14, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 20, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, 8192, 8191)
}
