package main

import "testing"

func runSample(t *testing.T, n int, expect string) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %s, but got %s", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, "34")
}

func TestSample2(t *testing.T) {
	runSample(t, 1, "1")
}
