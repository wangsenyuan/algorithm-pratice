package main

import "testing"

func runSample(t *testing.T, n int, expect int64) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 1260)
}
