package main

import "testing"

func runSample(t *testing.T, n int, expect string) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, alice)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, alice)
}

func TestSample3(t *testing.T) {
	runSample(t, 69, bob)
}
