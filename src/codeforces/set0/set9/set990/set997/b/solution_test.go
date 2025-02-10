package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 244)
}
