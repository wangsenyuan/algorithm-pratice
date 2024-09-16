package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 12, 124, 24, 4
	//
	runSample(t, "124", 4)
}

func TestSample2(t *testing.T) {
	// 12, 124, 24, 4
	//
	runSample(t, "04", 3)
}

func TestSample3(t *testing.T) {
	// 12, 124, 24, 4
	//
	runSample(t, "5810438174", 9)
}
