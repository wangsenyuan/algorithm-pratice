package main

import "testing"

func runSample(t *testing.T, n int, k int, m int, expect int) {
	res := solve(n, k, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, m := 3, 2, 1000000007
	expect := 4
	runSample(t, n, k, m, expect)
}

func TestSample2(t *testing.T) {
	n, k, m := 3, 4, 1000000007
	expect := 28
	runSample(t, n, k, m, expect)
}
