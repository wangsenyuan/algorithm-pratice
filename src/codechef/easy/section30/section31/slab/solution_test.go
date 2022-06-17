package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 600000
	expect := 577500
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 250000
	expect := 250000
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 1500000
	expect := 1312500
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 2500000
	expect := 2012500
	runSample(t, n, expect)
}
