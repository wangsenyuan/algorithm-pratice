package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	expect := 1
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	expect := 11
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 69228
	expect := 778304278
	runSample(t, n, expect)
}
