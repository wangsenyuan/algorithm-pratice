package main

import "testing"

func runSample(t *testing.T, n int, x int, pos int, expect int) {
	res := solve(n, x, pos)

	if res != expect {
		t.Fatalf("Sample expect %d,but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, p := 4, 1, 2
	expect := 6
	runSample(t, n, x, p, expect)
}

func TestSample2(t *testing.T) {
	n, x, p := 123, 42, 24
	expect := 824071958
	runSample(t, n, x, p, expect)
}
