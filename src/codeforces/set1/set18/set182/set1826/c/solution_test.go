package main

import "testing"

func runSample(t *testing.T, n int, m int, expect bool) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 95, 71
	expect := false
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 11, 67
	expect := false
	runSample(t, n, m, expect)
}
