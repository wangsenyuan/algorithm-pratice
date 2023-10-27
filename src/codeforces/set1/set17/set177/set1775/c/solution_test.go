package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 10, 8
	expect := 12
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n, x := 10, 10
	expect := 10
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n, x := 10, 42
	expect := -1
	runSample(t, n, x, expect)
}

func TestSample4(t *testing.T) {
	n, x := 20, 16
	expect := 24
	runSample(t, n, x, expect)
}

func TestSample5(t *testing.T) {
	n, x := 1000000000000000000, 0
	expect := 1152921504606846976
	runSample(t, n, x, expect)
}
