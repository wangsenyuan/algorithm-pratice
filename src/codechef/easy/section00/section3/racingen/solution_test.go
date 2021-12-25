package main

import "testing"

func runSample(t *testing.T, x, r, m int, expect bool) {
	res := solve(x, r, m)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 60, 3, 5, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 60, 3, 4, false)
}
