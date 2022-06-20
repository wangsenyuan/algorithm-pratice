package main

import "testing"

func runSample(t *testing.T, n int, x int64, expect bool) {
	res := solve(n, x)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 6, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 10, false)
}
