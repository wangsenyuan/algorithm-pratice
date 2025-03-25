package main

import "testing"

func runSample(t *testing.T, n int, k int, p int, expect int) {
	res := solve(n, k, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, p := 3, 1, 282282277
	expect := 16
	runSample(t, n, k, p, expect)
}

func TestSample2(t *testing.T) {
	n, k, p := 50, 25, 998244353
	expect := 131276976
	runSample(t, n, k, p, expect)
}
