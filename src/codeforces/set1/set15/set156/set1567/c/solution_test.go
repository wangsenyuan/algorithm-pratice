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
	expect := 9
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 12
	expect := 4
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	expect := 7
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 2021
	expect := 44
	runSample(t, n, expect)
}

func TestSample5(t *testing.T) {
	n := 10000
	expect := 99
	runSample(t, n, expect)
}
