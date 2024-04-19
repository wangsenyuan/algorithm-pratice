package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, 14)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 3, 14)
}

func TestSample4(t *testing.T) {
	runSample(t, 506, 2708, 3218)
}
