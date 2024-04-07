package main

import "testing"

func runSample(t *testing.T, r, b, p int, expect int) {
	res := solve(r, b, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 2, 63)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 3, 5, 3264)
}

func TestSample4(t *testing.T) {
	runSample(t, 6, 2, 9, 813023575)
}
