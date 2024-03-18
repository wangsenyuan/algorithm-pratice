package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 998244353
	expect := 5
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 42, 998244353
	expect := 793019428
	runSample(t, n, m, expect)
}
