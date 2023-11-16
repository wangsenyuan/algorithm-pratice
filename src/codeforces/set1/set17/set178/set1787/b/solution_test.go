package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 100
	expect := 20
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	expect := 10
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 864
	expect := 22
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 130056192
	expect := 118
	runSample(t, n, expect)
}

func TestSample5(t *testing.T) {
	n := 1000000000
	expect := 90
	runSample(t, n, expect)
}

func TestSample6(t *testing.T) {
	n := 999999018
	expect := 333333009
	runSample(t, n, expect)
}
