package main

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 12, 9)
}
