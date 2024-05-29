package main

import "testing"

func runSample(t *testing.T, k int, x int, a int, expect bool) {
	res := solve(k, x, a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 6, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 7, true)
}
