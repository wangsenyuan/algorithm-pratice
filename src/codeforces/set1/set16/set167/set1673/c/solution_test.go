package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, 74)
}

func TestSample3(t *testing.T) {
	runSample(t, 101, 8275126)
}
