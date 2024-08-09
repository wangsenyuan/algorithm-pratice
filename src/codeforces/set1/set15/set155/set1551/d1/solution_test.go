package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect bool) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 0, true)
}
