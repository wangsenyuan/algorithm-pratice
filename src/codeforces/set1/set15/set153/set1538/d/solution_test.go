package main

import "testing"

func runSample(t *testing.T, a int, b int, k int, expect bool) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 36, 48, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000000000, 1000000000, 1000000000, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2, 1, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 2, 1, true)
}
