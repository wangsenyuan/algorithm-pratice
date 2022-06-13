package main

import "testing"

func runSample(t *testing.T, n int, x int, expect bool) {
	res := solve(n, x)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 2, 3
	expect := false
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n, x := 2, 4
	expect := true
	runSample(t, n, x, expect)
}
