package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	expect := 748683265
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	expect := 748683265
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	expect := 842268673
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 200000
	expect := 202370013
	runSample(t, n, expect)
}
