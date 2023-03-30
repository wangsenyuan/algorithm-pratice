package main

import "testing"

func runSample(t *testing.T, n int, m int, l int, r int, expect int) {
	res := solve(n, m, l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, l, r := 5, 5, 1, 5
	expect := 12310
	runSample(t, n, m, l, r, expect)
}

func TestSample2(t *testing.T) {
	n, m, l, r := 5, 5, 4, 5
	expect := 4202
	runSample(t, n, m, l, r, expect)
}
