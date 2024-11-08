package main

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := solve(k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 21, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 0)
}
