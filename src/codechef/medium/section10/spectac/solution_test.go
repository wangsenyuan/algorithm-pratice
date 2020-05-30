package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, mod int, expect int) {
	res := solve(n, m, k, mod)

	if res != expect {
		t.Errorf("Sample %d %d %d %d, expect %d, but got %d", n, m, k, mod, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k, mod := 2, 2, 2, 997
	expect := 1
	runSample(t, n, m, k, mod, expect)
}

func TestSample2(t *testing.T) {
	n, m, k, mod := 30, 25, 20, 997
	expect := 687
	runSample(t, n, m, k, mod, expect)
}
