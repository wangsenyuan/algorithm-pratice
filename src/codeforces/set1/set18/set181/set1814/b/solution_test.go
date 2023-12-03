package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 6, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 4, 6)
}
