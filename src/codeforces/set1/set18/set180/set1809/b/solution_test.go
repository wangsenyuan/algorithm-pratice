package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 975461057789971042, 987654321)
}
