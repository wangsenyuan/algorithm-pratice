package main

import "testing"

func runSample(t *testing.T, x, y string, expect bool) {
	res := solve([]byte(x), []byte(y))

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "rrrjj", "rrrj", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "rj", "jr", false)
}
