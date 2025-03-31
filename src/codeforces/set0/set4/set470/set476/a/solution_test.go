package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 29, 7
	// n = a + 2 * b
	// a + b = m * w
	// b = n - m * w
	expect := 21
	runSample(t, n, m, expect)
}
