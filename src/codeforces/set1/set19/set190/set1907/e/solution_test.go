package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 9999999, 1522435234375)
}
