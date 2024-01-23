package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 179, 27)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000000000000000000, 2648956421)
}

func TestSample3(t *testing.T) {
	runSample(t, 465620822584281989, 1834608237)
}
