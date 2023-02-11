package main

import "testing"

func runSample(t *testing.T, x uint64, n int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	var x uint64 = 1
	expect := -1
	runSample(t, x, n, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	var x uint64 = 2
	expect := 4
	runSample(t, x, n, expect)
}

func TestSample3(t *testing.T) {
	n := 13
	var x uint64 = 42
	expect := 12
	runSample(t, x, n, expect)
}

func TestSample4(t *testing.T) {
	n := 19
	var x uint64 = 75804550
	expect := 11
	runSample(t, x, n, expect)
}
