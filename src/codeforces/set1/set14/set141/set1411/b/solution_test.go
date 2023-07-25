package main

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 282, 288)
}

func TestSample2(t *testing.T) {
	runSample(t, 1234567890, 1234568040)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000000000, 1000000000000000000)
}
