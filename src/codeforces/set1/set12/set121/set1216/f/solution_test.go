package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	s := "00100"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 1
	s := "000000"
	expect := 21
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 1
	s := "0011"
	expect := 4
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 12, 6
	s := "000010000100"
	expect := 15
	runSample(t, n, k, s, expect)
}
