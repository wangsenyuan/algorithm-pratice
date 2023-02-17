package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 256, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 512, false)
}
