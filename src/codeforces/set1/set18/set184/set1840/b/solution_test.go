package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 179, 100
	expect := 180
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 2
	expect := 4
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 2, 2
	expect := 3
	runSample(t, n, k, expect)
}
